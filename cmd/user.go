package cmd

import (
	"fmt"
	"syscall"

	"github.com/edznux/wonderxss/config"
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

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the application",
	Long: `
	Login to the application.
	It is useful to use the command line interface.
	`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		user := args[0]
		fmt.Println("Please enter a password:")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal("Could not read password", err)
		}
		password := string(bytePassword)

		fmt.Println("Please enter your OTP:")
		byteOTP, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal("Could not read OTP", err)
		}
		otp := string(byteOTP)

		log.Debugln("trying to connect the user", user)
		res, err := currentAPI.Login(user, password, otp)

		if err != nil {
			fmt.Println(err)
			//exit early
			return
		}
		config.SaveClientConfig(config.Client{
			Token:   res,
			Version: "v1",
		})
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
	rootCmd.AddCommand(loginCmd)
	userCmd.AddCommand(createUserCmd)
}
