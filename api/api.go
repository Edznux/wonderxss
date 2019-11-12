package api

import "github.com/edznux/wonder-xss/storage"

var store storage.Storage

func init() {
	store = storage.GetDB()
}
