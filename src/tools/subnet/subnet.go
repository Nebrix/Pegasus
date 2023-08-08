package main

import (
	"fmt"
	"os"
	"net"
	"math/big"
	"strconv"
)

func subnetCalculator(ipAddress string, cidr int) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		fmt.Println("Invalid IP address")
		return
	}

	ipInt := big.NewInt(0).SetBytes(ip.To4())
	ones, bits := cidr, 32
	mask := big.NewInt(0).Sub(big.NewInt(0).Lsh(big.NewInt(1), uint(bits-ones)), big.NewInt(1))

	networkInt := big.NewInt(0).And(ipInt, mask)
	broadcastInt := big.NewInt(0).Or(networkInt, big.NewInt(0).Not(mask))

	networkIP := net.IP(networkInt.Bytes())
	broadcastIP := net.IP(broadcastInt.Bytes())

	availableAddresses := new(big.Int).Sub(broadcastInt, networkInt)
	availableAddresses.Add(availableAddresses, big.NewInt(1))

	subnetCount := new(big.Int).Lsh(big.NewInt(1), uint(bits-cidr))

	fmt.Println("Subnet Details:")
	fmt.Printf("IP address: %s\n", ip)
	fmt.Printf("CIDR: /%d\n", cidr)
	fmt.Printf("Network address: %s\n", networkIP)
	fmt.Printf("Broadcast address: %s\n", broadcastIP)
	fmt.Printf("Number of available addresses: %s\n", availableAddresses)
	fmt.Printf("Number of subnets: %s\n", subnetCount)
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	ipAddress := os.Args[1]
	cidrStr := os.Args[2]

	cidr, err := strconv.Atoi(cidrStr)
	if err != nil {
		fmt.Println("Invalid CIDR:", err)
		return
	}

	subnetCalculator(ipAddress, cidr)
}
