package main

import (
	"github.com/edznux/wonderxss/cmd"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage"
)

func main() {

	// This populate the global config for the application
	config.Setup()
	storage.LoadStorageBackends()
	// This dispatch all the command from the application
	// It might be a `serve` command, or a create user command
	cmd.Execute()
}
