package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"

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

		path := args[0]
		name := filepath.Base(path)
		fmt.Printf("Adding payload: %s (%s)\n", name, path)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		}

		p, err := currentAPI.AddPayload(name, string(content), "application/javascript")
		if err != nil {
			log.Fatal("Could not create payload ", err)
		}
		fmt.Printf("Payload created: [%s] %s\n", p.ID, p.Name)
	},
}

// createCmd represents the create command
var getPayloadCmd = &cobra.Command{
	Use:   "get [payload]",
	Short: "Get all payloads or a specific one",
	Long:  `Get all payloads or a specific one by specifying it's it as a second argument`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			payloades, err := currentAPI.GetPayloads()
			if err != nil {
				log.Fatal("Could not get payloades", err)
			}
			fmt.Printf("%+v\n", payloades)
			return
		}

		payloadID := args[0]
		payload, err := currentAPI.GetPayload(payloadID)
		if err != nil {
			log.Fatal("Could not get payload"+payloadID, err)
		}
		fmt.Printf("%+v\n", payload)
		fmt.Printf("Payload: [%s] %s\n", payload.ID, payload.Name)

	},
}

func init() {
	rootCmd.AddCommand(payloadCmd)
	payloadCmd.AddCommand(createPayloadCmd)
	payloadCmd.AddCommand(getPayloadCmd)
}
