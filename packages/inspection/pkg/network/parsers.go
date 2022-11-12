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
