package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

const (
	md5Size    = 16 * 2
	sha1Size   = 20 * 2
	sha256Size = 32 * 2
	sha512Size = 64 * 2
)

func HashSha256(data string) {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	hash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("%v\n", hash)
}

func HashSha1(data string) {
	hasher := sha1.New()
	hasher.Write([]byte(data))
	hash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("%v\n", hash)
}

func HashSha512(data string) {
	hasher := sha512.New()
	hasher.Write([]byte(data))
	hash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("%v\n", hash)
}

func HashMD5(data string) {
	hasher := md5.New()
	hasher.Write([]byte(data))
	hash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("%v\n", hash)
}

func Decodehash(data string) string {
	switch len(data) {
	case md5Size:
		return "MD5"
	case sha1Size:
		return "SHA-1"
	case sha256Size:
		return "SHA-256"
	case sha512Size:
		return "SHA-512"
	default:
		return "Unknown"
	}
}
