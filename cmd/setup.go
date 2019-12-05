package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setupCmd represents the user command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setup start.")
		// This will init the storages methods, may be multiples

		db.Setup()
		fmt.Println("Setup done.")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
