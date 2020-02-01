package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/edznux/wonderxss/api"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// payloadCmd represents the payload command
var payloadCmd = &cobra.Command{
	Use:   "payload",
	Short: "Do all the operations on payloads.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			payloads, err := currentAPI.GetPayloads()
			if err != nil {
				log.Println(err)
			}
			tablePayloads(payloads)
			return
		}
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
		var payloads []api.Payload
		var err error

		if len(args) == 0 {
			payloads, err = currentAPI.GetPayloads()
			if err != nil {
				log.Fatal("Could not get payloads", err)
			}
			return
		} else {
			payloadID := args[0]
			payload, err := currentAPI.GetPayload(payloadID)
			if err != nil {
				log.Fatal("Could not get payload"+payloadID, err)
			}
			payloads = append(payloads, payload)
		}
		// TODO: replace the error from the api to custom api.Error
		// so we can do if err == api.ErrNotFound
		if payloads[0].ID != "" {
			tablePayloads(payloads)
		} else {
			fmt.Println("No payloads found.")
		}
	},
}

func tablePayloads(payloads []api.Payload) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Content", "Content Type", "Created At"})
	for _, p := range payloads {
		table.Append([]string{p.ID, p.Name, p.Content, p.CreatedAt.String()})
	}
	table.Render()
}

func init() {
	rootCmd.AddCommand(payloadCmd)
	payloadCmd.AddCommand(createPayloadCmd)
	payloadCmd.AddCommand(getPayloadCmd)
}
