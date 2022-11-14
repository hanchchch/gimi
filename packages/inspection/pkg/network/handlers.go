package network

import (
	"net/http"

	"github.com/google/gopacket"
)

func HttpHandler(callback func(*http.Request)) func(gopacket.Packet) {
	return func(packet gopacket.Packet) {
		if req := ParseHTTP(packet); req != nil {
			callback(req)
		}
	}
}

func DnsQueryHandler(callback func(string)) func(gopacket.Packet) {
	return func(packet gopacket.Packet) {
		if host := ParseDnsQuery(packet); host != nil {
			callback(string(*host))
		}
	}
}
