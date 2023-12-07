package sniffing

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func printPacketInfo(packet gopacket.Packet) {
	networkLayer := packet.NetworkLayer()
	transportLayer := packet.TransportLayer()
	applicationLayer := packet.ApplicationLayer()

	fmt.Println("Packet Captured:")
	fmt.Printf("  Timestamp: %v\n", packet.Metadata().Timestamp)
	fmt.Printf("     Length: %v bytes\n\n", packet.Metadata().CaptureLength)

	if networkLayer != nil {
		fmt.Printf("  Network Layer: %v\n", networkLayer.LayerType())
		fmt.Printf("      Source IP: %v\n", networkLayer.NetworkFlow().Src())
		fmt.Printf("        Dest IP: %v\n\n", networkLayer.NetworkFlow().Dst())
	}

	if transportLayer != nil {
		fmt.Printf("  Transport Layer: %v\n", transportLayer.LayerType())
		fmt.Printf("      Source Port: %v\n", transportLayer.TransportFlow().Src())
		fmt.Printf("        Dest Port: %v\n\n", transportLayer.TransportFlow().Dst())
	}

	if applicationLayer != nil {
		fmt.Printf("  Application Layer: %v\n", applicationLayer.LayerType())
		fmt.Printf("            Payload: %v\n\n", applicationLayer.Payload())
	}
}

func PacketSniffer(networkInterface string) {
	handle, err := pcap.OpenLive(networkInterface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		printPacketInfo(packet)
	}
}
