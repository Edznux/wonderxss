package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/edznux/wonderxss/api"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	defaultPayloadHeader = []string{"ID", "Name", "Content", "Content Type", "Created At"}
	fieldsPayload        []string
)

// payloadCmd represents the payload command
var payloadCmd = &cobra.Command{
	Use:     "payload",
	Aliases: []string{"payloads"},
	Short:   "Do all the operations on payloads.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			payloads, err := currentAPI.GetPayloads()
			if err != nil {
				log.Errorln(err)
				return
			}
			renderPayloads(payloads)
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
		} else {
			payloadID := args[0]
			payload, err := currentAPI.GetPayload(payloadID)
			if err != nil {
				log.Fatal("Could not get payload "+payloadID, err)
			}
			payloads = append(payloads, payload)
		}
		renderPayloads(payloads)
	},
}

// deletePayloadsCmd represents the delete command
var deletePayloadsCmd = &cobra.Command{
	Use:   "delete [payload]",
	Short: "delete an payload",
	Long:  `delete an payload based on it's ID`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		ID := args[0]
		err := currentAPI.DeletePayload(ID)
		if err != nil {
			log.Fatal("Could not delete payload ", err)
		}
		fmt.Printf("Payload deleted: [%s] \n", ID)
	},
}

func renderPayloads(payloads []api.Payload) {
	rows := buildPayloadsTable(payloads)
	if isRaw {
		renderRaw(rows)
		return
	}

	if len(payloads) > 0 {
		renderTable(rows)
	} else {
		fmt.Println("No payloads found.")
	}
}

func buildPayloadsTable(payloads []api.Payload) [][]string {
	var rows [][]string
	rows = make([][]string, len(payloads))

	for i, p := range payloads {
		content := ""
		rows[i] = make([]string, 0)
		for _, f := range fields {
			if f == "Content" {
				if isReplace {
					content = p.Content
				} else {
					content = replacePlaceholders(p.Content)
				}
				rows[i] = append(rows[i], content)
			} else {
				rows[i] = append(rows[i], getFieldString(p, f))
			}
		}
	}
	return rows
}

func init() {
	rootCmd.AddCommand(payloadCmd)
	payloadCmd.AddCommand(createPayloadCmd)
	payloadCmd.AddCommand(getPayloadCmd)
	payloadCmd.AddCommand(deletePayloadsCmd)
	payloadCmd.PersistentFlags().StringSliceVar(&fieldsPayload, "fields", defaultPayloadHeader, "Fields you want to query")
}
