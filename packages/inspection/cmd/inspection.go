package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hanchchch/gimi/packages/inspection/pkg/aws"
	c "github.com/hanchchch/gimi/packages/inspection/pkg/chrome"
	h "github.com/hanchchch/gimi/packages/inspection/pkg/headless"
	n "github.com/hanchchch/gimi/packages/inspection/pkg/network"
	"github.com/hanchchch/gimi/packages/inspection/pkg/urls"
	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/proto"
)

const (
	bound = "-----"
)

func inspect(url string, device string, chromeArgs []string) (*pb.InspectionResult, c.ChromeInspectResult) {
	r := &pb.InspectionResult{
		Url:       url,
		Malicious: false,
		Reasons:   []string{},
		Hosts:     []string{},
		SendingTo: []string{},
		Urls:      []string{},
	}

	hc := h.NewHeadlessClient()
	cc, err := c.NewChromeClient(c.ChromeClientOptions{ChromeArgs: chromeArgs})
	if err != nil {
		panic(err)
	}

	ni := n.NewNetworkInspector(device)
	ni.AppendHandler(n.HttpHandler(func(req *http.Request) {
		url := req.URL.String()
		if strings.HasPrefix(url, "/") {
			url = req.Host + url
		}
		log.Printf("captured - http - %v %v", req.Method, url)
		r.Urls = append(r.Urls, url)
	}))
	ni.AppendHandler(n.DnsQueryHandler(func(host string) {
		log.Printf("captured - dns - %v", host)
		r.Hosts = append(r.Hosts, host)
	}))
	go ni.Listen()
	defer ni.Terminate()

	log.Printf("visiting with simple http client")
	if hr, err := hc.Visit(h.VisitParams{
		Method: "GET",
		Url:    url,
	}); err != nil {
		panic(err)
	} else {
		r.Locations = hr.Locations
	}

	log.Printf("starting chrome crawler")
	cr, err := cc.Run(url)
	if err != nil {
		panic(err)
	}

	r.SendingTo = cr.SendingTo
	if cr.Reasons != nil {
		r.Malicious = len(cr.Reasons) > 0
		r.Reasons = cr.Reasons
	}

	return r, cr
}

func uploadScreenshot(url string, screenshot []byte, s3region string, s3bucket string, accessKey string, secretKey string) string {
	if s3, err := aws.NewS3Client(aws.S3ClientOptions{
		AccessKeyId:     accessKey,
		SecretAccessKey: secretKey,
		Region:          s3region,
	}); err != nil {
		panic(err)
	} else {
		location, err := s3.UploadScreenshot(s3bucket, url, screenshot)
		log.Printf("uploaded screenshot - %v", location)
		if err != nil {
			panic(err)
		}
		return location
	}
}

func output(r *pb.InspectionResult) {
	log.Printf("result - %v", r.String())
	b, err := proto.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v%v", bound, string(b))
}

func main() {
	godotenv.Load()

	url := flag.String("url", "", "target url")
	device := flag.String("device", "", "network interface")
	ua := flag.String("user-agent", "", "user agent")
	s3region := flag.String("s3-region", "ap-northeast-2", "s3 region")
	s3bucket := flag.String("s3-bucket", "gimi-screenshots", "s3 bucket")

	flag.Parse()

	log.Printf("starting inspection for %v", *url)

	if *device == "" {
		*device = os.Getenv("NETWORK_INTERFACE")
	}
	log.Printf("device: %v", *device)

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	chromeArgs := []string{
		"disable-gpu",
		"disable-dev-shm-usage",
		"window-size=1920,1080",
		"no-sandbox",
		"headless",
	}

	if *ua != "" {
		log.Printf("user agent: %v", *ua)
		chromeArgs = append(chromeArgs, "user-agent="+*ua)
	}
	log.Printf("chrome args: %v", chromeArgs)

	log.Printf("inspecting...")
	r, cr := inspect(urls.EnsureProtocol(*url), *device, chromeArgs)

	log.Printf("uploading screenshot...")
	r.Screenshot = uploadScreenshot(*url, cr.Screenshot, *s3region, *s3bucket, awsAccessKey, awsSecretKey)

	output(r)
}
