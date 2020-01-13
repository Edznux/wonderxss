package local

import "github.com/edznux/wonderxss/storage"

type Local struct {
	store storage.Storage
}

func New() *Local {
	return &Local{store: storage.GetDB()}
}

func (local *Local) GetHealth() (string, error) {
	return "OK", nil
}
