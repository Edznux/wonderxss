package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// aliasesCmd represents the injection command
var aliasesCmd = &cobra.Command{
	Use:   "alias",
	Short: "Do all the operations on aliases.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Invalid arguments, see `Available Commands`")
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

var getAliasesCmd = &cobra.Command{
	Use:   "get [alias]",
	Short: "Create a new alias",
	Long: `Create a new alias by providing a file path to the alias.
	The (basename of the) path will be the name of the alias`,
	Run: func(cmd *cobra.Command, args []string) {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Alias", "Created At"})

		if len(args) == 0 {
			aliases, err := currentAPI.GetAliases()
			if err != nil {
				log.Fatal("Could not get aliases", err)
			}
			for _, a := range aliases {
				table.Append([]string{a.ID, a.Alias, a.CreatedAt.String()})
			}
			table.Render()
			return
		}

		aliasID := args[0]
		alias, err := currentAPI.GetAlias(aliasID)
		if err != nil {
			fmt.Println("Could not get alias "+aliasID, err)
		}

		if alias.ID != "" {
			table.Append([]string{alias.ID, alias.Alias, alias.CreatedAt.String()})
			table.Render()
		} else {
			fmt.Println("No alias found")
		}

	},
}

func init() {
	rootCmd.AddCommand(aliasesCmd)
	aliasesCmd.AddCommand(createAliasesCmd)
	aliasesCmd.AddCommand(getAliasesCmd)
}
