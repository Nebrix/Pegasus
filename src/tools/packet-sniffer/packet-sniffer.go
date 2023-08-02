package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func packetHandler(packet gopacket.Packet) {
	// Replace this with the packet processing logic you want
	fmt.Println(packet)
}

func getDefaultNetworkDevice() string {
	// Find all available network devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Return the name of the first network device (the default one)
	if len(devices) > 0 {
		return devices[0].Name
	}

	log.Fatal("No network devices found")
	return ""
}

func main() {
	// Get the default network device
	defaultDevice := getDefaultNetworkDevice()

	// Open the default network device for packet capture
	handle, err := pcap.OpenLive(defaultDevice, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	fmt.Println("Starting packet sniffer on device:", defaultDevice)

	// Start processing packets
	for packet := range packetSource.Packets() {
		packetHandler(packet)
		
		time.Sleep(1 * time.Second)
	}
}
