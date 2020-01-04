package crypto

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"hash"

	"github.com/edznux/wonderxss/storage/models"
)

func Hash(data string, hashType string) []byte {
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

	return hash.Sum(nil)
}

func GenerateSRIHashes(data string) models.SRIHashes {
	return models.SRIHashes{
		SHA256: "sha256-" + base64.StdEncoding.EncodeToString(Hash(data, "sha256")),
		SHA384: "sha384-" + base64.StdEncoding.EncodeToString(Hash(data, "sha384")),
		SHA512: "sha512-" + base64.StdEncoding.EncodeToString(Hash(data, "sha512")),
	}
}

//GenerateOTPSecret generates the 80-bit base32 encoded string user's secret
func GenerateOTPSecret() (string, error) {
	secretLen := 16
	secret := make([]byte, secretLen)
	_, err := rand.Read(secret)
	if err != nil {
		fmt.Println("error:", err)
		return "", fmt.Errorf("Could not generate a new TOTP Secret")
	}
	hashed := sha1.Sum(secret)
	// return string(hashed[:]), nil
	return base32.StdEncoding.EncodeToString(hashed[:]), nil
}
