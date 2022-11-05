package main

import (
	"encoding/json"
	"flag"
	"fmt"

	c "github.com/hanchchch/gimi/packages/inspection/pkg/chrome"
	h "github.com/hanchchch/gimi/packages/inspection/pkg/headless"
)

type InspectionResult struct {
	Locations []string `json:"locations"`
	Malicious bool     `json:"malicious"`
}

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

	r := InspectionResult{}
	r.Locations = hr.Locations
	r.Malicious = cr.Malicious

	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(b))
}
