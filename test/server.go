package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	serverMessageIdentifier = "[SERVER]: "
	clientMessageIdentifier = "[CLIENT]: "
)

var (
	clients     = make(map[net.Conn]string)
	chatHistory []string
)

func handleClient(client net.Conn) {
	reader := bufio.NewReader(client)
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error receiving username:", err)
		client.Close()
		delete(clients, client)
		return
	}

	username = strings.TrimSpace(username)
	if username == "" {
		fmt.Println("Error: Username cannot be empty.")
		client.Close()
		delete(clients, client)
		return
	}

	clients[client] = username

	sendChatHistory(client)

	message := serverMessageIdentifier + "Welcome, " + username + "!"
	broadcast(message)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		message = strings.TrimSpace(message)
		if message != "" {
			if message == "server status" {
				sendStatus(client)
			} else {
				msg := username + ": " + message
				fmt.Println(msg)
				chatHistory = append(chatHistory, msg)
				broadcast(msg)
			}
		}
	}

	client.Close()
	delete(clients, client)
}

func broadcast(message string) {
	for client := range clients {
		_, err := client.Write([]byte(message + "\n"))
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}

func sendChatHistory(client net.Conn) {
	if len(chatHistory) == 0 {
		return
	}

	for _, msg := range chatHistory {
		_, err := client.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("Error sending chat history:", err)
			client.Close()
			delete(clients, client)
			return
		}
		time.Sleep(50 * time.Millisecond) // Small delay to avoid message merging
	}
}

func sendStatus(client net.Conn) {
	numMembers := len(clients)
	statusMessage := fmt.Sprintf(serverMessageIdentifier+"Number of members: %d\n", numMembers)
	_, err := client.Write([]byte(statusMessage))
	if err != nil {
		fmt.Println("Error sending status:", err)
	}
}

func main() {
	HOST := "127.0.0.1"
	PORT := "12345"

	l, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error creating server:", err)
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Chatroom server started.")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}
