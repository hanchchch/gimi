package network

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type NetworkInspector struct {
	device   string
	handlers []func(packet gopacket.Packet)
	handle   *pcap.Handle
}

func NewNetworkInspector(device string) *NetworkInspector {
	return &NetworkInspector{
		device:   device,
		handlers: []func(packet gopacket.Packet){},
	}
}

func (i *NetworkInspector) AppendHandler(f func(packet gopacket.Packet)) {
	i.handlers = append(i.handlers, f)
}

func (i *NetworkInspector) Listen() {
	if handle, err := pcap.OpenLive(i.device, 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else {
		i.handle = handle
	}

	packetSource := gopacket.NewPacketSource(i.handle, i.handle.LinkType())
	for packet := range packetSource.Packets() {
		for _, handler := range i.handlers {
			handler(packet)
		}
	}
}

func (i *NetworkInspector) Terminate() {
	i.handle.Close()
}
