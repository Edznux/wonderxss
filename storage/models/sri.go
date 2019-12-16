package models

import "encoding/json"

type SRIHashes struct {
	SHA256 string `json:"sha256"`
	SHA384 string `json:"sha384"`
	SHA512 string `json:"sha512"`
}

func (sh *SRIHashes) String() string {
	js, err := json.Marshal(sh)
	if err != nil {
		panic("Could not stringify the SRIHashes struct")
	}
	return string(js)
}
