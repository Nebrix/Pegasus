package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	domain := os.Args[1]

	// Perform DNS lookup
	records, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Print IP addresses
	fmt.Printf("IP addresses for %s:\n", domain)
	for _, ip := range records {
		fmt.Println(ip)
	}

	// Perform DNS lookup for other record types
	fmt.Printf("\nOther DNS records for %s:\n", domain)

	// NS records
	nsRecords, err := net.LookupNS(domain)
	if err == nil {
		for _, ns := range nsRecords {
			fmt.Printf("NS: %s\n", ns.Host)
		}
	}

	// MX records
	mxRecords, err := net.LookupMX(domain)
	if err == nil {
		for _, mx := range mxRecords {
			fmt.Printf("MX: %s, Preference: %d\n", mx.Host, mx.Pref)
		}
	}

	// SOA record (extracted from TXT records)
	txtRecords, err := net.LookupTXT(domain)
	if err == nil {
		for _, txt := range txtRecords {
			if strings.HasPrefix(txt, "v=spf1") {
				soaParts := strings.Fields(txt)
				if len(soaParts) >= 7 {
					fmt.Printf("SOA: Primary NS: %s, Responsible person's mailbox: %s\n", soaParts[6], soaParts[1])
				}
				break
			}
		}
	}

	// TXT records
	txtRecords, err = net.LookupTXT(domain)
	if err == nil {
		for _, txt := range txtRecords {
			fmt.Printf("TXT: %s\n", txt)
		}
	}
}
