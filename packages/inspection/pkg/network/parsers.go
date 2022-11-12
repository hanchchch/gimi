package network

import (
	"bufio"
	"bytes"
	"net/http"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func ParseHTTP(packet gopacket.Packet) *http.Request {
	if packet.ApplicationLayer() == nil || packet.TransportLayer() == nil || packet.NetworkLayer() == nil {
		return nil
	}

	if packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
		return nil
	}
	if len(packet.TransportLayer().LayerPayload()) == 0 {
		return nil
	}

	reader := bufio.NewReader(bytes.NewReader(packet.TransportLayer().LayerPayload()))
	req, err := http.ReadRequest(reader)
	if err != nil {
		return nil
	}

	return req
}

func ParseDnsQuery(packet gopacket.Packet) *string {
	if packet.ApplicationLayer() == nil || packet.TransportLayer() == nil || packet.NetworkLayer() == nil {
		return nil
	}

	if packet.ApplicationLayer().LayerType() == layers.LayerTypeDNS {
		dns := packet.ApplicationLayer().LayerContents()
		// OR OPCODE
		//  0   0000
		if (dns[2] | 0x07) != 0x07 {
			return nil
		}
		if len(dns) <= 12 {
			return nil
		}
		ptr := 12
		query := []byte{}
		for int(dns[ptr]) != 0 {
			query = append(query, dns[ptr:ptr+int(dns[ptr])+1]...)
			query = append(query, '.')
			ptr += int(dns[ptr]) + 1
		}
		result := string(query[:len(query)-1])
		return &result
	}

	return nil
}
