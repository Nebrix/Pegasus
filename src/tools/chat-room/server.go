package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var clients = make(map[net.Conn]bool)
var mutex sync.Mutex

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat room server started on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		mutex.Lock()
		clients[conn] = true
		mutex.Unlock()

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	clientAddr := conn.RemoteAddr().String()

	fmt.Printf("%s joined the chat\n", clientAddr)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		message = strings.TrimSpace(message)

		if message == "/exit" {
			break
		}

		broadcastMessage(fmt.Sprintf("%s: %s", clientAddr, message))
	}

	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()

	fmt.Printf("%s left the chat\n", clientAddr)
}

func broadcastMessage(message string) {
	mutex.Lock()
	defer mutex.Unlock()

	for client := range clients {
		_, err := client.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error broadcasting message:", err)
			mutex.Lock()
			delete(clients, client)
			mutex.Unlock()
		}
	}
}
