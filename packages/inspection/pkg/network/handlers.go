package network

import (
	"github.com/google/gopacket"
)

func HttpHostHandler(callback func(string)) func(gopacket.Packet) {
	return func(packet gopacket.Packet) {
		if req := ParseHTTP(packet); req != nil {
			callback(string(req.Host))
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
