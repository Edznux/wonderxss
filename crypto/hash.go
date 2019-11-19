package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
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

type SRIHashes struct {
	SHA256 string `json:"sha256"`
	SHA384 string `json:"sha384"`
	SHA512 string `json:"sha512"`
}

func GenerateSRIHashes(data string) SRIHashes {
	return SRIHashes{
		SHA256: "sha256-" + base64.StdEncoding.EncodeToString(Hash(data, "sha256")),
		SHA384: "sha384-" + base64.StdEncoding.EncodeToString(Hash(data, "sha384")),
		SHA512: "sha512-" + base64.StdEncoding.EncodeToString(Hash(data, "sha512")),
	}
}

func (sh *SRIHashes) String() string {
	js, err := json.Marshal(sh)
	if err != nil {
		panic("Could not stringify the SRIHashes struct")
	}
	return string(js)
}
