package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanPort(host string, port int, timeout time.Duration, results chan<- int) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err == nil {
		conn.Close()
		results <- port
	}
}

func scanPorts(host string, startPort, endPort int, timeout time.Duration) []int {
	results := make(chan int)
	openPorts := []int{}

	for port := startPort; port <= endPort; port++ {
		go scanPort(host, port, timeout, results)
	}

	timeoutCh := time.After(timeout)

	for i := startPort; i <= endPort; i++ {
		select {
		case port := <-results:
			openPorts = append(openPorts, port)
		case <-timeoutCh:
			break
		}
	}

	close(results)
	return openPorts
}

func main() {
	if len(os.Args) < 4 {
		return
	}

	host := os.Args[1]
	startPort, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid startPort:", err)
		return
	}

	endPort, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid endPort:", err)
		return
	}

	timeout := 2 * time.Second

	openPorts := scanPorts(host, startPort, endPort, timeout)

	fmt.Printf("Open ports on %s:\n", host)
	for _, port := range openPorts {
		fmt.Printf("%d\n", port)
	}
}
