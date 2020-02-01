package main

import (
	"github.com/edznux/wonderxss/cmd"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage"
	"github.com/sirupsen/logrus"
)

func main() {

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	// This populate the global config for the application
	config.Setup()
	storage.LoadStorageBackends()
	// This dispatch all the command from the application
	// It might be a `serve` command, or a create user command
	cmd.Execute()
}
