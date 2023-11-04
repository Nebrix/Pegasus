package tools

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

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
