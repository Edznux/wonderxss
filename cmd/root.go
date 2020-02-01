package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/api/http/client"
	"github.com/edznux/wonderxss/api/local"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage"
	"github.com/spf13/cobra"
)

var (
	db         storage.Storage
	remote     bool
	currentAPI api.API
)

var rootCmd = &cobra.Command{
	Use:   "wonderxss",
	Short: "WonderXSS is a pentest tool for discovering Blind XSSs",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if remote {
			log.Debugln("Using remote API!")
			cfg, err := config.ReadClientConfig()
			if err != nil {
				log.Fatalln("Coulnd read client config:", err.Error())
			}
			currentAPI = client.New(cfg)
		} else {
			log.Debugln("Using local API")
			currentAPI = local.New()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// No arguments called. Abort and print help
		cmd.Help()
	},
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Get the health of the application",
	Run: func(cmd *cobra.Command, args []string) {
		// No arguments called. Abort and print help
		result, err := currentAPI.GetHealth()
		if err != nil {
			fmt.Println("Error :", err.Error())
		}
		fmt.Println(result)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&remote, "remote", "t", false, "remote bool")
	rootCmd.AddCommand(healthCmd)
}
