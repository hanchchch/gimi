package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "target url")

	flag.Parse()

	fmt.Printf("url: %s", *url)
}
