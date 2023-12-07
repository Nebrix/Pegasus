package osint

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/likexian/whois"
)

func GetDNSRecords(host string) {
	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Println("Error resolving A records:", err)
	} else {
		fmt.Printf("A (IPv4) addresses for %v:\n", host)
		for _, addr := range addrs {
			fmt.Println(addr)
		}
	}

	nsRecords, err := net.LookupNS(host)
	if err != nil {
		fmt.Println("Error resolving NS records:", err)
	} else {
		fmt.Printf("NS (Name Server) records for %v:\n", host)
		for _, ns := range nsRecords {
			fmt.Println(ns.Host)
		}
	}

	cname, err := net.LookupCNAME(host)
	if err != nil {
		fmt.Println("Error resolving CNAME records:", err)
	} else {
		fmt.Printf("CNAME (Canonical Name) for %v: %v\n", host, cname)
	}

	txtRecords, err := net.LookupTXT(host)
	if err != nil {
		fmt.Println("Error resolving TXT records:", err)
	} else {
		fmt.Printf("TXT (Text) records for %v:\n", host)
		for _, txt := range txtRecords {
			fmt.Println(txt)
		}
	}
}

func getInternalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", nil
}

func getPublicIP() (string, error) {
	resp, err := http.Get("https://ifconfig.co/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	publicIP := strings.TrimSpace(string(body))
	return publicIP, nil
}

func ShowIP() {
	internalIP, err := getInternalIP()
	if err != nil {
		fmt.Println("Error getting internal IP:", err)
	} else {
		fmt.Println("Internal IP:", internalIP)
	}

	publicIP, err := getPublicIP()
	if err != nil {
		fmt.Println("Error getting public IP:", err)
	} else {
		fmt.Println("Public IP:", publicIP)
	}
}

const (
	hostName = "ipinfo.io"
	port     = "80"
)

func GetIpInfo(host string) {
	path := "/" + host + "/json"
	url := "http://" + hostName + ":" + port + path

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP request failed with status code: %v\n", resp.StatusCode)
		os.Exit(1)
	}

	buffer := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			fmt.Print(string(buffer[:n]))
		}
		if err != nil {
			break
		}
	}
}

func Getwhois(host string) {
	result, err := whois.Whois(host)
	if err == nil {
		fmt.Println(result)
	}
}

func getWebsiteHeaders(url string) (http.Header, error) {
	response, err := http.Head(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return response.Header, nil
}

func HeaderRetrieve(url string) {
	headers, err := getWebsiteHeaders(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Headers for %s:\n", url)
	for key, values := range headers {
		fmt.Printf("%s: %v\n", key, values)
	}
}
