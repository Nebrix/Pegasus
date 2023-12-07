package cmd

import (
	"fmt"
	"shell/src/helper"
	"shell/src/network"
	"shell/src/network/scanning"
	"shell/src/network/sniffing"
	"shell/src/osint"
	"strconv"
	"time"
)

func Traceroute(arguments []string) {
	network.Traceroute(arguments[0])
}

func Webheader(arguments []string) {
	osint.HeaderRetrieve(arguments[0])
}

func PortScanner(arguments []string) {
	scanning.PortScanner(arguments[0])
}

func Ping(arguments []string) {
	host, count := arguments[0], 0
	if len(arguments) > 1 {
		count, _ = strconv.Atoi(arguments[1])
	}

	network.Ping(host, count, 2*time.Second)
}

func Whois(arguments []string) {
	osint.Getwhois(arguments[0])
}

func Hash(arguments []string) {
	method, data := arguments[0], arguments[1]

	switch method {
	case "md5", "sha1", "sha256", "sha512":
		network.Hash(method, data)
	case "decode":
		fmt.Printf("Hash Algorithm: %s\n", network.DecodeHash(data))
	default:
		helper.HandleWarn("Unsupported hash method", method)
	}
}

func ShowIP() {
	osint.ShowIP()
}

func ShowSubnet(arguments []string) {
	if cidr, err := strconv.Atoi(arguments[1]); !helper.HandleErr("Invalid CIDR", err) {
		network.ShowCalculation(arguments[0], cidr)
	}
}

func GetDNSRecords(arguments []string) {
	osint.GetDNSRecords(arguments[0])
}

func GetIPInfo(arguments []string) {
	osint.GetIpInfo(arguments[0])
}

func ShowSnifferPackets(arguments []string) {
	sniffing.PacketSniffer(arguments[0])
}
