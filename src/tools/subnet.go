package tools

import (
	"fmt"
	"net"
	"os"
)

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
	fmt.Printf("IP address: %s\n", ipAddress)
	fmt.Printf("CIDR: /%d\n", cidr)
	fmt.Printf("Network address: %s\n", networkIP.String())
	fmt.Printf("Broadcast address: %s\n", broadcastIP.String())
	fmt.Printf("Number of available addresses: %d\n", availableAddresses)
	fmt.Printf("Number of subnets: %d\n", subnetCount)
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
