package tools

import (
	"fmt"
	"net/http"
	"os"
)

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
