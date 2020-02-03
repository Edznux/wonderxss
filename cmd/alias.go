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

// aliasesCmd represents the injection command
var aliasesCmd = &cobra.Command{
	Use:     "alias",
	Aliases: []string{"aliases"},
	Short:   "Do all the operations on aliases.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			aliases, err := currentAPI.GetAliases()
			if err != nil {
				log.Println(err)
			}
			if len(aliases) > 0 {
				tableAliases(aliases)
			} else {
				fmt.Println("No Aliases found.")
			}
			return
		}
		cmd.Help()
	},
}

// createAliasesCmd represents the create command
var createAliasesCmd = &cobra.Command{
	Use:   "create [alias]",
	Short: "Create a new alias",
	Long: `Create a new alias by providing a file path to the alias.
	The (basename of the) path will be the name of the alias`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		path := args[0]
		name := filepath.Base(path)
		fmt.Printf("Adding alias: %s (%s)\n", name, path)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		}

		p, err := currentAPI.AddAlias(name, string(content))
		if err != nil {
			log.Fatal("Could not create alias ", err)
		}
		fmt.Printf("Alias created: [%s] %s\n", p.ID, p.Alias)
	},
}

// deleteAliasesCmd represents the delete command
var deleteAliasesCmd = &cobra.Command{
	Use:   "delete [alias]",
	Short: "delete an alias",
	Long:  `delete an alias based on it's ID`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		ID := args[0]
		err := currentAPI.DeleteAlias(ID)
		if err != nil {
			log.Fatal("Could not delete alias ", err)
		}
		fmt.Printf("Alias deleted: [%s] \n", ID)
	},
}

var getAliasesCmd = &cobra.Command{
	Use:   "get [alias]",
	Short: "Create a new alias",
	Long: `Create a new alias by providing a file path to the alias.
	The (basename of the) path will be the name of the alias`,
	Run: func(cmd *cobra.Command, args []string) {
		var aliases []api.Alias
		var err error
		if len(args) == 0 {
			aliases, err = currentAPI.GetAliases()
			if err != nil {
				log.Fatal("Could not get aliases", err)
			}
		} else {
			aliasID := args[0]
			alias, err := currentAPI.GetAlias(aliasID)
			if err != nil {
				fmt.Println("Could not get alias "+aliasID, err)
			}
			aliases = append(aliases, alias)
		}
		if len(aliases) > 0 {
			tableAliases(aliases)
		} else {
			fmt.Println("No aliases found.")
		}
	},
}

func tableAliases(payloads []api.Alias) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Alias", "Created At"})
	for _, p := range payloads {
		table.Append([]string{p.ID, p.Alias, p.CreatedAt.String()})
	}
	table.Render()
}

func init() {
	rootCmd.AddCommand(aliasesCmd)
	aliasesCmd.AddCommand(createAliasesCmd)
	aliasesCmd.AddCommand(getAliasesCmd)
	aliasesCmd.AddCommand(deleteAliasesCmd)
}
