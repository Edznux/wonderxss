package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// injectionCmd represents the injection command
var injectionCmd = &cobra.Command{
	Use:   "injection",
	Short: "Do all the operations on injections.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Invalid arguments, see `Available Commands`")
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

func init() {
	rootCmd.AddCommand(injectionCmd)
	injectionCmd.AddCommand(createInjectionCmd)
}
