package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 8000

	// Define a file server that serves files from the current directory
	fileServer := http.FileServer(http.Dir("."))

	// Create a custom handler that adds CORS headers
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
		fileServer.ServeHTTP(w, r)
	})

	// Start the server
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Serving at %s\n", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
