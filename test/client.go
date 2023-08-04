package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	serverMessageIdentifier = "[SERVER]: "
	clientMessageIdentifier = "[CLIENT]: "
	promptText              = "Enter Message: "
)

var lastMessage string
var mutex = make(chan struct{}, 1)

func receiveMessages(clientSocket net.Conn) {
	reader := bufio.NewReader(clientSocket)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error receiving message:", err)
			break
		}

		if message != "" {
			// Check if the message is not empty (to avoid printing empty lines)
			if !startsWithAnyIdentifier(message) {
				mutex <- struct{}{} // Acquire the mutex
				clearLastMessage()
				fmt.Print(message)
				fmt.Print(promptText)
				lastMessage = message
				<-mutex // Release the mutex
			}
		}
	}
}

func startsWithAnyIdentifier(message string) bool {
	identifiers := []string{serverMessageIdentifier, clientMessageIdentifier}
	for _, id := range identifiers {
		if strings.HasPrefix(message, id) {
			return true
		}
	}
	return false
}

func sendMessages(clientSocket net.Conn, username string) {
	reader := make(chan string)
	go func() {
		for {
			message := readInputLine()
			if message != "" {
				clientSocket.Write([]byte(clientMessageIdentifier + message + "\n"))
			}
			reader <- message
		}
		close(reader)
	}()

	for message := range reader {
		mutex <- struct{}{} // Acquire the mutex
		clearLastMessage()
		if message != "" {
			fmt.Print(promptText)
		}
		fmt.Print(message)
		lastMessage = message
		<-mutex // Release the mutex
	}
}

func readInputLine() string {
	fmt.Print(promptText)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()
	return message
}

func clearLastMessage() {
	if lastMessage != "" {
		lines := strings.Count(lastMessage, "\n") + 1
		fmt.Printf("\033[%dA", lines) // Move cursor up to previous lines
		fmt.Print("\033[K")           // Clear the line
	}
}

func main() {
	HOST := "127.0.0.1"
	PORT := "12345"

	fmt.Print("Enter your username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()

	if strings.TrimSpace(username) == "" {
		fmt.Println("Error: Username cannot be empty.")
		return
	}

	clientSocket, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer clientSocket.Close()

	_, err = clientSocket.Write([]byte(username + "\n"))
	if err != nil {
		fmt.Println("Error sending username:", err)
		return
	}

	fmt.Println("Connected to server. Start typing your messages.")
	fmt.Print(promptText)

	go receiveMessages(clientSocket)
	sendMessages(clientSocket, username)
}
