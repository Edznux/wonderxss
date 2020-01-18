package local

import (
	"log"

	"github.com/edznux/wonderxss/storage"
)

type Local struct {
	store storage.Storage
}

func New() *Local {
	log.Println("New Local API")
	str := storage.GetDB()
	return &Local{store: str}
}

func (local *Local) GetHealth() (string, error) {
	return "OK", nil
}
