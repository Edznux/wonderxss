package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/edznux/wonderxss/api"
	"github.com/spf13/cobra"
)

// payloadCmd represents the payload command
var payloadCmd = &cobra.Command{
	Use:   "payload",
	Short: "Do all the operations on payloads.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Invalid arguments, see `Available Commands`")
		cmd.Help()
	},
}

// createCmd represents the create command
var createPayloadCmd = &cobra.Command{
	Use:   "create [payload]",
	Short: "Create a new payload",
	Long: `Create a new payload by providing a file path to the payload.
	The (basename of the) path will be the name of the payload`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		api.Init()
		path := args[0]
		name := filepath.Base(path)
		fmt.Printf("Adding payload: %s (%s)\n", name, path)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		}

		p, err := api.AddPayload(name, string(content))
		if err != nil {
			log.Fatal("Could not create payload ", err)
		}
		fmt.Printf("Payload created: [%s] %s\n", p.ID, p.Name)
	},
}

func init() {
	rootCmd.AddCommand(payloadCmd)
	payloadCmd.AddCommand(createPayloadCmd)
}
