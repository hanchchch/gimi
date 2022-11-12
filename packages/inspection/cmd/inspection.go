package main

import (
	"flag"
	"fmt"
	"os"

	c "github.com/hanchchch/gimi/packages/inspection/pkg/chrome"
	h "github.com/hanchchch/gimi/packages/inspection/pkg/headless"
	n "github.com/hanchchch/gimi/packages/inspection/pkg/network"
	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
	"google.golang.org/protobuf/proto"
)

func main() {
	url := flag.String("url", "", "target url")
	device := flag.String("device", "", "network interface")
	ua := flag.String("user-agent", "", "user agent")

	flag.Parse()

	if *device == "" {
		*device = os.Getenv("NETWORK_INTERFACE")
	}

	r := &pb.InspectionResult{
		Url:   *url,
		Hosts: []string{},
	}

	ni := n.NewNetworkInspector(*device)
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
		Url:    *url,
	})
	if err != nil {
		panic(err)
	}

	r.Locations = hr.Locations

	chromeArgs := []string{
		"disable-gpu",
		"window-size=1920,1080",
		"no-sandbox",
		"headless",
	}

	if *ua != "" {
		chromeArgs = append(chromeArgs, "user-agent="+*ua)
	}

	cc, err := c.NewChromeClient(c.ChromeClientOptions{ChromeArgs: chromeArgs})
	if err != nil {
		panic(err)
	}

	cr, err := cc.Run(*url)
	if err != nil {
		panic(err)
	}

	r.Malicious = cr.Malicious

	b, err := proto.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", string(b))
}
