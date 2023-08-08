package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type IPData struct {
	City    string `json:"city"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Org     string `json:"org"`
	Loc     string `json:"loc"`
	Postal  string `json:"postal"`
}

func ipLookup(ipAddress string) (*IPData, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s/json", ipAddress)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("IP lookup request failed with status code: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data IPData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	ipAddress := os.Args[1]
	data, err := ipLookup(ipAddress)
	if err != nil {
		fmt.Println("IP lookup failed:", err)
		return
	}

	fmt.Printf("Location: %s, %s, %s\n", data.City, data.Region, data.Country)
	fmt.Printf("ISP: %s\n", data.Org)
	fmt.Printf("Location: %s\n", data.Loc)
	fmt.Printf("Postal Code: %s\n", data.Postal)
}
