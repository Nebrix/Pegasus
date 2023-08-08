package main

import (
	"fmt"
	"os"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
)

func identifyHash(hashValue string) string {
	switch len(hashValue) {
	case md5.Size * 2:
		return "MD5"
	case sha1.Size * 2:
		return "SHA-1"
	case sha256.Size * 2:
		return "SHA-256"
	default:
		return "Unknown"
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	hashValue := os.Args[1]

	hashAlgorithm := identifyHash(hashValue)
	fmt.Println("Hash Algorithm:", hashAlgorithm)
}
