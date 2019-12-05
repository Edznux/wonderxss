package cmd

import (
	"fmt"
	"os"

	"github.com/edznux/wonderxss/storage"
	"github.com/spf13/cobra"
)

var db storage.Storage

var rootCmd = &cobra.Command{
	Use:   "wonderxss",
	Short: "WonderXSS is a pentest tool for discovering Blind XSSs",
	Run: func(cmd *cobra.Command, args []string) {
		// No arguments called. Abort and print help
		cmd.Help()
	},
}

func Execute() {
	// set the db globaly. It's already initialised, just "globalify" the object...
	db = storage.GetDB()
	db.Init()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
