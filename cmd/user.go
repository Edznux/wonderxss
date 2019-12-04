package cmd

import (
	"fmt"
	"log"

	"github.com/edznux/wonderxss/api"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Do all the operations on users.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Invalid arguments, see `Available Commands`")
		cmd.Help()
	},
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [username]",
	Short: "Create a new user",
	Long:  `Create a new user by providing a username and filling out informations`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		user := args[0]
		fmt.Println("create called", user)
		password := "test"
		u, err := api.CreateUser(user, password)
		if err != nil {
			log.Fatal("Could not create user")
		}
		fmt.Printf("User created: [%s] %s", u.ID, u.Username)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(createCmd)
}
