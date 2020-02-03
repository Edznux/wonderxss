package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/edznux/wonderxss/api"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

const (
	REPLACE_URL_TAG     = "##URL##"
	REPLACE_SRI_TAG     = "##SRI_HASH##"
	REPLACE_CROSSORIGIN = "##CROSSORIGIN##"
)

var (
	SRIKinds     = []string{"sha256", "sha384", "sha512"}
	CROSSORIGIN  = []string{"anonymous", "use-credentials"}
	url          string
	sri          string
	crossorigin  string
	payload      string
	protocol     string
	isRaw        bool
	useSubdomain bool
	useHTTPS     bool
)

// injectionCmd represents the injection command
var injectionCmd = &cobra.Command{
	Aliases: []string{"injection", "injections"},
	Use:     "injection",
	Short:   "Do all the operations on injections.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		globalFlagsSetup(cmd, args)
		protocol = "http://"
		if useHTTPS {
			protocol = "https://"
		}
		//FIXME Change port (handle https with 443)
		if useSubdomain {
			url = protocol + payload + "." + cfg.Host
		} else {
			url = protocol + cfg.Host + "/p/" + payload
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			injections, err := currentAPI.GetInjections()
			if err != nil {
				log.Println(err)
			}
			if len(injections) > 0 {
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
		log.Debugf("Current api ?%+v", currentAPI)
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
		if len(injections) > 0 {
			tableInjections(injections)
		} else {
			fmt.Println("No injection found.")
		}
	},
}

// deleteInjectionsCmd represents the delete command
var deleteInjectionsCmd = &cobra.Command{
	Use:   "delete [injection]",
	Short: "delete an injection",
	Long:  `delete an injection based on it's ID`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		ID := args[0]
		err := currentAPI.DeleteInjection(ID)
		if err != nil {
			log.Fatal("Could not delete injection ", err)
		}
		fmt.Printf("Injection deleted: [%s] \n", ID)
	},
}

func tableInjections(injections []api.Injection) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Content", "Created At"})
	var content string
	for _, p := range injections {
		if isRaw {
			content = p.Content
		} else {
			content = replacePlaceholders(p.Content)
		}
		table.Append([]string{p.ID, p.Name, content, p.CreatedAt.String()})
	}
	table.Render()
}

func replacePlaceholders(input string) string {
	input = strings.Replace(input, REPLACE_URL_TAG, url, -1)
	input = strings.Replace(input, REPLACE_SRI_TAG, sri, -1)
	input = strings.Replace(input, REPLACE_CROSSORIGIN, crossorigin, -1)
	return input
}

func init() {

	rootCmd.AddCommand(injectionCmd)

	injectionCmd.AddCommand(createInjectionCmd)
	injectionCmd.AddCommand(getInjectionCmd)
	injectionCmd.AddCommand(deleteInjectionsCmd)

	injectionCmd.PersistentFlags().BoolVar(&isRaw, "raw", false, "Do not replace placeholder in the injections")
	injectionCmd.PersistentFlags().BoolVar(&useSubdomain, "use-subdomain", true, "Use the subdomain as the payload id (enabled by default)")
	injectionCmd.PersistentFlags().BoolVar(&useHTTPS, "use-https", true, "Use HTTPS (enabled by default)")
	injectionCmd.PersistentFlags().StringVar(&sri, "sri", "sha256", "SRI Type [sha256,sha384,sha512]")
	injectionCmd.PersistentFlags().StringVar(&payload, "payload", "", "Payload ID or Name")
}
