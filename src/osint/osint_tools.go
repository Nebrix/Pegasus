package osint

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	areacodes "shell/src/osint/areaCodes"
	"strings"

	"github.com/likexian/whois"
)

const (
	EXIT_SUCCESS = 0
	EXIT_FAILURE = 1
	BUFFER_LEN   = 16
	PN_STR_LEN   = 12
)

const (
	TITLE_SPLASH_STR = `    ____  _                      ____  _   _ ____
   |  _ \| |__   ___  _ __   ___|  _ \| \ | / ___|
   | |_) | '_ \ / _ \| '_ \ / _ \ | | |  \| \___ \
   |  __/| | | | (_) | | | |  __/ |_| | |\  |___) |
   |_|   |_| |_|\___/|_| |_|\___|____/|_| \_|____/
  `

	INFO_SPLASH_STR = `
  § PhoneDNS Nebrix | Licensed under MIT License
  § Original creator: https://github.com/Nebrix
  § Creator of this version: https://github.com/Nebrix`
)

func strLen(buffer string) int {
	return len(buffer)
}

func isDigit(ch byte) bool {
	return bool(ch >= '0' && ch <= '9')
}

func isDigitsAndDashes(pnStr string) bool {
	isDigits := true
	firstDash := 3
	secondDash := 7

	for i := 0; i < PN_STR_LEN-1; i++ {
		if i == firstDash || i == secondDash {
			continue
		}
		if !isDigit(pnStr[i]) {
			isDigits = false
			break
		}
	}
	return isDigits
}

func isPhoneNumber(buffer string) bool {
	correctLen := strLen(buffer) == PN_STR_LEN
	correctDash := buffer[3] == '-' && buffer[7] == '-'
	correctDigits := isDigitsAndDashes(buffer)
	return correctLen && correctDash && correctDigits
}

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

	fmt.Printf("Headers for %v:\n", url)
	for key, values := range headers {
		fmt.Printf("%v: %v\n", key, values)
	}
}

func PhoneDNS() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(TITLE_SPLASH_STR)
	fmt.Println(INFO_SPLASH_STR)

	for {
		fmt.Println("------------------")
		fmt.Printf("DNS -> ")
		scanner.Scan()
		inputBuffer := scanner.Text()

		if inputBuffer == "exit" {
			return
		}

		isValidNumber := isPhoneNumber(string(inputBuffer))
		areaCodeIndex := areacodes.GetAreaCode(string(inputBuffer))
		hasValidAreaCode := areaCodeIndex < len(areacodes.AreaCodes) && areacodes.AreaCodes[areaCodeIndex].State != ""

		if !hasValidAreaCode {
			fmt.Fprintln(os.Stderr, "Invalid area code.")
			isValidNumber = false
		}

		if isValidNumber && hasValidAreaCode {
			fmt.Printf("State    :%v\n", areacodes.AreaCodes[areaCodeIndex].State)
			fmt.Printf("Country  :%v\n", areacodes.AreaCodes[areaCodeIndex].Country)
			fmt.Printf("Timezone :%v\n", areacodes.AreaCodes[areaCodeIndex].Timezone)
			fmt.Printf("Region   :%v\n", areacodes.AreaCodes[areaCodeIndex].Regions)
		} else {
			fmt.Fprintln(os.Stderr, "Invalid phone number format. Please enter a valid phone number in the format xxx-xxx-xxxx.")
		}
	}
}
