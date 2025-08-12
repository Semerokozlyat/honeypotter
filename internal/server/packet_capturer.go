package server

import (
	"fmt"
	"io"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"

	"github.com/Semerokozlyat/honeypotter/internal/config"
)

type PacketCapturer struct {
	handleTCP    *pcap.Handle
	packetSrcTCP *gopacket.PacketSource
}

func NewPacketCapturer(cfg *config.PacketCapturer) (*PacketCapturer, error) {
	handle, err := pcap.OpenLive(cfg.InterfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		return nil, fmt.Errorf("open handle for network interface %s: %w", cfg.InterfaceName, err)
	}
	// Capture only TCP traffic
	err = handle.SetBPFFilter("tcp")
	if err != nil {
		return nil, fmt.Errorf("set bpf filter for handle: %w", err)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	return &PacketCapturer{
		handleTCP:    handle,
		packetSrcTCP: packetSource,
	}, nil
}

func (pc *PacketCapturer) Run() error {
	log.Println("running packet capturer")
	for {
		packet, err := pc.packetSrcTCP.NextPacket()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("failed to fetch next packet from stream:", err)
			continue
		}
		err = pc.handleTCPPacket(packet)
		if err != nil {
			log.Println("failed to handle packet:", err)
			continue
		}
	}
	return nil
}

func (pc *PacketCapturer) handleTCPPacket(p gopacket.Packet) error {
	tcpLayer := p.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		log.Println("no TCP level in packet, next")
		return nil
	}
	tcp, _ := tcpLayer.(*layers.TCP)
	fmt.Printf("TCP Packet: %s:%d -> %s:%d (Seq: %d)\n",
		p.NetworkLayer().NetworkFlow().Src().String(),
		tcp.SrcPort,
		p.NetworkLayer().NetworkFlow().Dst().String(),
		tcp.DstPort,
		tcp.Seq)
	// TODO: add further analysis of TCP payloads, flags, etc.
	return nil
}

func (pc *PacketCapturer) Close() {
	pc.handleTCP.Close()
}
