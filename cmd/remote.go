package cmd

import (
	"fmt"
	"log"
	"syscall"

	"github.com/edznux/wonderxss/api/http/client"
	"github.com/edznux/wonderxss/config"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var c *client.Client
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Do all the operations on remote server.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ReadClientConfig()
		if err != nil {
			log.Fatalln(err)
		}
		c = client.New(cfg)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var remoteLogin = &cobra.Command{
	Use:   "login",
	Short: "Login on the remote server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Missing arguments users")
			cmd.Help()
			return
		}
		username := args[0]
		fmt.Println("Please enter a password:")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal("Could not read password", err)
		}
		password := string(bytePassword)
		//FIXME add the user for OTP AFTER getting the error server side.
		token, err := c.Login(username, password, "")
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(token)
	},
}
var remoteHealth = &cobra.Command{
	Use:   "health",
	Short: "Get the health of the remote server",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := c.GetHealth()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
	remoteCmd.AddCommand(remoteHealth)
	remoteCmd.AddCommand(remoteLogin)
}
