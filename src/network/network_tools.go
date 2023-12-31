package network

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	md5Size    = 16 * 2
	sha1Size   = 20 * 2
	sha256Size = 32 * 2
	sha512Size = 64 * 2
)

func Ping(host string, count int, timeout time.Duration) {
	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	conn, err := net.DialIP("ip4:icmp", nil, addr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	for i := 1; i <= count; i++ {
		msg := make([]byte, 64)
		msg[0] = 8
		msg[1] = 0
		msg[2] = 0
		msg[3] = 0
		msg[4] = 0
		msg[5] = 0
		msg[6] = 0
		msg[7] = byte(i)
		checksum := checkSum(msg)
		msg[2] = byte(checksum >> 8)
		msg[3] = byte(checksum & 0xFF)

		startTime := time.Now()

		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		reply := make([]byte, 64)
		conn.SetReadDeadline(time.Now().Add(timeout))
		_, err = conn.Read(reply)
		if err != nil {
			fmt.Println("Request timed out")
		} else {
			duration := time.Since(startTime)
			fmt.Printf("Reply from %v: time=%v\n", addr.String(), duration)
		}

		time.Sleep(1 * time.Second)
	}
}

func checkSum(msg []byte) uint16 {
	sum := uint32(0)
	for i := 0; i < len(msg); i += 2 {
		sum += uint32(msg[i+1]) | (uint32(msg[i]) << 8)
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)
}

func Hash(method, data string) {
	var hasher hash.Hash

	switch method {
	case "md5":
		hasher = md5.New()
	case "sha1":
		hasher = sha1.New()
	case "sha256":
		hasher = sha256.New()
	case "sha512":
		hasher = sha512.New()
	default:
		fmt.Printf("Unsupported hash method: %s\n", method)
		return
	}

	hasher.Write([]byte(data))
	hash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("%v\n", hash)
}

func DecodeHash(data string) string {
	switch len(data) {
	case md5Size:
		return "MD5"
	case sha1Size:
		return "SHA-1"
	case sha256Size:
		return "SHA-256"
	case sha512Size:
		return "SHA-512"
	default:
		return "Unknown"
	}
}

func subnetCalculator(ipAddress string, cidr int) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		fmt.Println("Invalid IP address")
		return
	}

	if cidr < 0 || cidr > 32 {
		fmt.Println("Invalid CIDR")
		return
	}

	ipInt := ipToUint32(ip)
	ones, bits := cidr, 32
	mask := (1<<uint32(bits-ones) - 1)

	networkInt := ipInt & ^uint32(mask)
	broadcastInt := ipInt | uint32(mask)

	networkIP := uint32ToIP(networkInt)
	broadcastIP := uint32ToIP(broadcastInt)

	availableAddresses := broadcastInt - networkInt + 1
	subnetCount := uint32(1 << uint32(bits-cidr))

	fmt.Println("Subnet Details:")
	fmt.Printf("IP address: %v\n", ipAddress)
	fmt.Printf("CIDR: /%v\n", cidr)
	fmt.Printf("Network address: %v\n", networkIP.String())
	fmt.Printf("Broadcast address: %v\n", broadcastIP.String())
	fmt.Printf("Number of available addresses: %v\n", availableAddresses)
	fmt.Printf("Number of subnets: %v\n", subnetCount)
}

func ipToUint32(ip net.IP) uint32 {
	ip = ip.To4()
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

func uint32ToIP(ipInt uint32) net.IP {
	return net.IPv4(byte(ipInt>>24), byte(ipInt>>16), byte(ipInt>>8), byte(ipInt))
}

func ShowCalculation(host string, cidr int) {
	if cidr < 0 || cidr > 32 {
		fmt.Println("Invalid CIDR")
		os.Exit(1)
	}

	subnetCalculator(host, cidr)
}

func Traceroute(host string) {
	fmt.Printf("Traceroute to %v...\n", host)

	ipAddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		fmt.Println("Error resolving target host:", err)
		os.Exit(1)
	}

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Println("Error opening scoket:", err)
		os.Exit(1)
	}
	defer conn.Close()

	ttl := 1

	for {
		message := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  ttl,
				Data: []byte(""),
			},
		}

		messageBytes, err := message.Marshal(nil)
		if err != nil {
			fmt.Println("Error marshalling ICMP message:", err)
			os.Exit(1)
		}

		if err := conn.IPv4PacketConn().SetTTL(ttl); err != nil {
			fmt.Println("Error setting TTL:", err)
			os.Exit(1)
		}

		_, err = conn.WriteTo(messageBytes, ipAddr)
		if err != nil {
			fmt.Println("Error sending ICMP message:", err)
			os.Exit(1)
		}

		reply := make([]byte, 1500)
		_, _, err = conn.ReadFrom(reply)
		if err != nil {
			fmt.Println("Error receiving ICMP reply:", err)
			os.Exit(1)
		}

		fmt.Printf("%d. %s\n", ttl, ipAddr.String())

		if ipAddr.String() == host {
			break
		}

		ttl++

		if ttl > 30 {
			fmt.Println("Max hops reached. Exiting.")
			break
		}
	}
}
