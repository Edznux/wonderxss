package cmd

import (
	"fmt"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
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
var createUserCmd = &cobra.Command{
	Use:   "create [username]",
	Short: "Create a new user",
	Long:  `Create a new user by providing a username and filling out informations`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		username := args[0]
		user, err := currentAPI.GetUserByName(username)
		if user.ID != "" {
			fmt.Println("User already exist.")
			return
		}

		fmt.Println("Please enter a password:")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal("Could not read password", err)
		}
		password := string(bytePassword)

		u, err := currentAPI.CreateUser(username, password)

		if err != nil {
			log.Fatal("Could not create user ", err)
		}
		fmt.Printf("User created: [%s] %s\n", u.ID, u.Username)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(createUserCmd)
}
