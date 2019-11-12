package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)

func Hash(data string, hashType string) string {
	fmt.Println("Hash function :", data, hashType)
	var hash hash.Hash
	if hashType == "sha256" {
		hash = sha256.New()
		hash.Write([]byte(data))
	}
	if hashType == "sha384" {
		hash = sha512.New384()
		hash.Write([]byte(data))
	}
	if hashType == "sha512" {
		hash = sha512.New()
		hash.Write([]byte(data))
	}

	return hex.EncodeToString(hash.Sum(nil))
}
