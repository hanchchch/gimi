package main

import (
	"flag"

	c "github.com/hanchchch/gimi/packages/inspection/pkg/chrome"
	h "github.com/hanchchch/gimi/packages/inspection/pkg/headless"
	pb "github.com/hanchchch/gimi/packages/proto/go/messages"
	"google.golang.org/protobuf/proto"
)

func main() {
	url := flag.String("url", "", "target url")

	flag.Parse()

	hr, err := h.NewHeadlessClient().Visit(h.VisitParams{
		Method: "GET",
		Url:    *url,
	})
	if err != nil {
		panic(err)
	}

	cc, err := c.NewChromeClient()
	if err != nil {
		panic(err)
	}

	cr, err := cc.Inspect(*url)
	if err != nil {
		panic(err)
	}

	r := &pb.InspectionResult{
		Url: *url,
	}
	r.Locations = hr.Locations
	r.Malicious = cr.Malicious

	b, err := proto.Marshal(r)
	if err != nil {
		panic(err)
	}

	println(string(b))
}
