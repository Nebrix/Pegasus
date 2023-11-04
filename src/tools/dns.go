package tools

import (
	"fmt"
	"net"
)

func GetDNSRecords(host string) {
	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Println("Error resolving A records:", err)
	} else {
		fmt.Printf("A (IPv4) addresses for %s:\n", host)
		for _, addr := range addrs {
			fmt.Println(addr)
		}
	}

	nsRecords, err := net.LookupNS(host)
	if err != nil {
		fmt.Println("Error resolving NS records:", err)
	} else {
		fmt.Printf("NS (Name Server) records for %s:\n", host)
		for _, ns := range nsRecords {
			fmt.Println(ns.Host)
		}
	}

	cname, err := net.LookupCNAME(host)
	if err != nil {
		fmt.Println("Error resolving CNAME records:", err)
	} else {
		fmt.Printf("CNAME (Canonical Name) for %s: %s\n", host, cname)
	}

	txtRecords, err := net.LookupTXT(host)
	if err != nil {
		fmt.Println("Error resolving TXT records:", err)
	} else {
		fmt.Printf("TXT (Text) records for %s:\n", host)
		for _, txt := range txtRecords {
			fmt.Println(txt)
		}
	}
}
