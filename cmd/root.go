package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/edznux/wonderxss/config"
	"github.com/spf13/cobra"
)

var cfg config.Config

func init() {
	var err error
	cfg, err = config.Load("")
	if err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "wonderxss",
	Short: "WonderXSS is a pentest tool for discovering Blind XSSs",
	Run: func(cmd *cobra.Command, args []string) {
		// No arguments called. Abort and print help
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
