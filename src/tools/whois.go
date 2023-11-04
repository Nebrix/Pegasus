package tools

import (
	"fmt"

	"github.com/likexian/whois"
)

func Getwhois(host string) {
	result, err := whois.Whois(host)
	if err == nil {
		fmt.Println(result)
	}
}
