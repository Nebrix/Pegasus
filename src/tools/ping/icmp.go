package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	icmpProtocol  = 1
	icmpEchoCode  = 8
	icmpEchoReply = 0
)

func calculateChecksum(msg []byte) uint16 {
	sum := 0
	for i := 0; i < len(msg)-1; i += 2 {
		sum += int(msg[i])<<8 | int(msg[i+1])
	}

	if len(msg)%2 != 0 {
		sum += int(msg[len(msg)-1]) << 8
	}

	sum = (sum >> 16) + (sum & 0xFFFF)
	sum = sum + (sum >> 16)
	return uint16(^sum)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <destination IP address>")
		os.Exit(1)
	}

	destAddrStr := os.Args[1]
	destAddr := net.ParseIP(destAddrStr)

	if destAddr == nil {
		fmt.Println("Invalid destination IP address.")
		os.Exit(1)
	}

	conn, err := net.DialIP("ip4:icmp", nil, &net.IPAddr{IP: destAddr})
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	const echoPayloadSize = 32
	echoPayload := make([]byte, echoPayloadSize)
	for i := range echoPayload {
		echoPayload[i] = byte(i)
	}

	const echoIdentifier = 1234
	const echoSequenceNumber = 1

	echoMsg := make([]byte, 8+echoPayloadSize)
	echoMsg[0] = icmpEchoCode
	echoMsg[1] = 0
	echoMsg[2] = 0
	echoMsg[3] = 0
	echoMsg[4] = byte(echoIdentifier >> 8)
	echoMsg[5] = byte(echoIdentifier & 0xFF)
	echoMsg[6] = byte(echoSequenceNumber >> 8)
	echoMsg[7] = byte(echoSequenceNumber & 0xFF)

	copy(echoMsg[8:], echoPayload)

	checksum := calculateChecksum(echoMsg)
	echoMsg[2] = byte(checksum >> 8)
	echoMsg[3] = byte(checksum & 0xFF)

	// Send the ICMP echo request
	startTime := time.Now()
	_, err = conn.Write(echoMsg)
	if err != nil {
		fmt.Println("Error sending ICMP request:", err)
		os.Exit(1)
	}

	// Set a deadline for reading the response
	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		fmt.Println("Error setting read deadline:", err)
		os.Exit(1)
	}

	// Receive the ICMP echo reply
	reply := make([]byte, 1500)
	n, _, err := conn.ReadFrom(reply)
	if err != nil {
		fmt.Println("Error receiving ICMP reply:", err)
		os.Exit(1)
	}

	// Calculate round-trip time
	rtt := time.Since(startTime)

	// Check if the received packet is an ICMP echo reply
	if reply[0] == icmpEchoReply && reply[4] == echoMsg[4] && reply[5] == echoMsg[5] {
		fmt.Printf("Received %d bytes from %s: icmp_seq=%d time=%v\n", n, destAddr, echoSequenceNumber, rtt)
	} else {
		fmt.Println("Received an invalid ICMP reply.")
	}

}