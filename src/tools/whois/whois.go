package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	domain := os.Args[1]

	// Connect to the WHOIS server
	conn, err := net.Dial("tcp", "whois.iana.org:43")
	if err != nil {
		fmt.Println("Error connecting to WHOIS server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Send the domain name to the WHOIS server
	fmt.Fprintf(conn, "%s\r\n", domain)

	// Read the response from the WHOIS server
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response from WHOIS server:", err)
		os.Exit(1)
	}
}
