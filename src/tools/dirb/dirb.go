package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func checkDirExists(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	defer response.Body.Close()

	// Check if the HTTP status code is 200 (OK) or 301 (Moved Permanently)
	if response.StatusCode == http.StatusOK || response.StatusCode == http.StatusMovedPermanently {
		return true
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	baseURL := os.Args[1]

	// Read the wordlist from the file
	file, err := os.Open("lists/*.txt")
	if err != nil {
		fmt.Println("Error opening wordlist.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var directories []string

	// Store each line (word) from the file in the directories slice
	for scanner.Scan() {
		directories = append(directories, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading wordlist.txt:", err)
		return
	}

	fmt.Println("Directory buster started...")

	// Iterate over the directories and make requests
	for _, dir := range directories {
		url := baseURL + "/" + dir
		if checkDirExists(url) {
			fmt.Printf("[%s] Directory found: %s\n", http.MethodGet, url)
		}
	}
}