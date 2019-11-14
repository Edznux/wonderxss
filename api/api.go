package api

import "github.com/edznux/wonderxss/storage"

var store storage.Storage

func InitApi() {
	store = storage.GetDB()
}
