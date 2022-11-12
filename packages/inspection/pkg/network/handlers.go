package network

import (
	"github.com/google/gopacket"
)

func HandleHttpHost(packet gopacket.Packet) {
	if req := ParseHTTP(packet); req != nil {
		// fmt.Printf("%v\n", string(req.Host))
		// TODO - do something with the http host
		return
	}
}
