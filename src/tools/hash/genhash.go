package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"os"
)

func hashString(s string, algorithm string) string {
	switch algorithm {
	case "md5":
		h := md5.New()
		h.Write([]byte(s))
		return fmt.Sprintf("%x", h.Sum(nil))
	case "sha1":
		h := sha1.New()
		h.Write([]byte(s))
		return fmt.Sprintf("%x", h.Sum(nil))
	case "sha256":
		h := sha256.New()
		h.Write([]byte(s))
		return fmt.Sprintf("%x", h.Sum(nil))
	default:
		return "Error: Unsupported algorithm"
	}
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	string := os.Args[1]
	algorithm := os.Args[2]

	fmt.Println(hashString(string, algorithm))
}
