package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hanchchch/gimi/packages/inspection/pkg/aws"
	c "github.com/hanchchch/gimi/packages/inspection/pkg/chrome"
	h "github.com/hanchchch/gimi/packages/inspection/pkg/headless"
	n "github.com/hanchchch/gimi/packages/inspection/pkg/network"
	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/proto"
)

const (
	s3region = "ap-northeast-2"
	s3bucket = "gimi-screenshots"
)

func inspect(url string, device string, chromeArgs []string) (*pb.InspectionResult, c.ChromeInspectResult) {
	r := &pb.InspectionResult{
		Url:   url,
		Hosts: []string{},
	}

	ni := n.NewNetworkInspector(device)
	ni.AppendHandler(n.HttpHostHandler(func(host string) {
		r.Hosts = append(r.Hosts, host)
	}))
	ni.AppendHandler(n.DnsQueryHandler(func(host string) {
		r.Hosts = append(r.Hosts, host)
	}))
	go ni.Listen()
	defer ni.Terminate()

	hr, err := h.NewHeadlessClient().Visit(h.VisitParams{
		Method: "GET",
		Url:    url,
	})
	if err != nil {
		panic(err)
	}

	r.Locations = hr.Locations

	cc, err := c.NewChromeClient(c.ChromeClientOptions{ChromeArgs: chromeArgs})
	if err != nil {
		panic(err)
	}

	cr, err := cc.Run(url)
	if err != nil {
		panic(err)
	}

	r.Malicious = cr.Malicious

	return r, cr
}

func uploadScreenshot(url string, screenshot []byte, accessKey string, secretKey string) {
	if s3, err := aws.NewS3Client(aws.S3ClientOptions{
		AccessKeyId:     accessKey,
		SecretAccessKey: secretKey,
		Region:          s3region,
	}); err != nil {
		panic(err)
	} else {
		if err := s3.UploadScreenshot(s3bucket, url, screenshot); err != nil {
			panic(err)
		}
	}
}

func output(r *pb.InspectionResult) {
	b, err := proto.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", string(b))
}

func main() {
	godotenv.Load()

	url := flag.String("url", "", "target url")
	device := flag.String("device", "", "network interface")
	ua := flag.String("user-agent", "", "user agent")

	flag.Parse()

	if *device == "" {
		*device = os.Getenv("NETWORK_INTERFACE")
	}

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	chromeArgs := []string{
		"disable-gpu",
		"window-size=1920,1080",
		"no-sandbox",
		"headless",
	}

	if *ua != "" {
		chromeArgs = append(chromeArgs, "user-agent="+*ua)
	}

	r, cr := inspect(*url, *device, chromeArgs)

	uploadScreenshot(*url, cr.Screenshot, awsAccessKey, awsSecretKey)

	output(r)
}
