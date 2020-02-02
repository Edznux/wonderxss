package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/edznux/wonderxss/api"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// injectionCmd represents the injection command
var injectionCmd = &cobra.Command{
	Use:   "injection",
	Short: "Do all the operations on injections.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			injections, err := currentAPI.GetInjections()
			if err != nil {
				log.Println(err)
			}
			// TODO: replace the error from the api to custom api.Error
			// so we can do if err == api.ErrNotFound
			if injections[0].ID != "" {
				tableInjections(injections)
			} else {
				fmt.Println("No injections found.")
			}
			return
		}
		cmd.Help()
	},
}

// createCmd represents the create command
var createInjectionCmd = &cobra.Command{
	Use:   "create [injection]",
	Short: "Create a new injection",
	Long: `Create a new injection by providing a file path to the injection.
	The (basename of the) path will be the name of the injection`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		path := args[0]
		name := filepath.Base(path)
		fmt.Printf("Adding injection: %s (%s)\n", name, path)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		}

		p, err := currentAPI.AddInjection(name, string(content))
		if err != nil {
			log.Fatal("Could not create injection ", err)
		}
		fmt.Printf("Injection created: [%s] %s\n", p.ID, p.Name)
	},
}

var getInjectionCmd = &cobra.Command{
	Use:   "get [injection]",
	Short: "Get all injections or a specific one",
	Long:  `Get all injections or a specific one by specifying it's it as a second argument`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var injections []api.Injection
		if len(args) == 0 {
			injections, err = currentAPI.GetInjections()
			if err != nil {
				log.Fatal("Could not get injections", err)
			}
		} else {
			injectionID := args[0]
			injection, err := currentAPI.GetInjection(injectionID)
			if err != nil {
				log.Warnf("Could not get Injection: [%s], %s", injectionID, err)
			}
			injections = append(injections, injection)
		}
		// TODO: replace the error from the api to custom api.Error
		// so we can do if err == api.ErrNotFound
		if injections[0].ID != "" {
			tableInjections(injections)
		} else {
			fmt.Println("No injection found.")
		}
	},
}

func tableInjections(injections []api.Injection) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Content", "Created At"})
	for _, p := range injections {
		table.Append([]string{p.ID, p.Name, p.Content, p.CreatedAt.String()})
	}
	table.Render()
}

func init() {
	rootCmd.AddCommand(injectionCmd)
	injectionCmd.AddCommand(createInjectionCmd)
	injectionCmd.AddCommand(getInjectionCmd)
}
