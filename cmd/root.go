package cmd

import (
	"fmt"
	"log"
	"os"

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
	fmt.Println(remote)
	// set the db globaly. It's already initialised, just "globalify" the object...
	if remote {
		fmt.Println("Using remote API!")
		cfg, err := config.ReadClientConfig()
		if err != nil {
			log.Fatalln("Coulnd read client config:", err.Error())
		}
		log.Printf("Config: %+v\n", cfg)
		currentAPI = client.New(cfg)
	} else {
		log.Println("Using local API")
		currentAPI = local.New()
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	var remoteLocal bool
	rootCmd.PersistentFlags().BoolVarP(&remoteLocal, "remote", "r", false, "use WonderXSS on remote host")
	fmt.Println(remoteLocal)
	rootCmd.AddCommand(healthCmd)
}
