package main

import (
	"encoding/json"
	"flag"
	"fmt"

	h "github.com/hanchchch/gimi/packages/inspection/pkg/headless"
)

type InspectionResult struct {
	Locations []string `json:"locations"`
}

func main() {
	url := flag.String("url", "", "target url")

	flag.Parse()

	vr, err := h.NewHeadlessClient().Visit(h.VisitParams{
		Method: "GET",
		Url:    *url,
	})

	if err != nil {
		panic(err)
	}

	r := InspectionResult{}
	r.Locations = vr.Locations

	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(b))
}
