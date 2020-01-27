package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

// aliasesCmd represents the injection command
var aliasesCmd = &cobra.Command{
	Use:   "aliases",
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
		fmt.Printf("Adding injection: %s (%s)\n", name, path)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		}

		p, err := currentAPI.AddAlias(name, string(content))
		if err != nil {
			log.Fatal("Could not create injection ", err)
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
		if len(args) == 0 {
			aliases, err := currentAPI.GetAliases()
			if err != nil {
				log.Fatal("Could not get aliases", err)
			}
			fmt.Printf("%+v\n", aliases)
			return
		}

		aliasID := args[0]
		alias, err := currentAPI.GetAlias(aliasID)
		if err != nil {
			log.Fatal("Could not get alias"+aliasID, err)
		}
		fmt.Printf("%+v\n", alias)
		fmt.Printf("Alias: [%s] %s\n", alias.ID, alias.Alias)
	},
}

func init() {
	rootCmd.AddCommand(aliasesCmd)
	aliasesCmd.AddCommand(createAliasesCmd)
	aliasesCmd.AddCommand(getAliasesCmd)
}
