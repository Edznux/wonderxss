package api

import "github.com/edznux/wonderxss/storage"

var store storage.Storage

func Init() {
	store = storage.GetDB()
}
