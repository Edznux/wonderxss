package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/edznux/wonderxss/api"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

const (
	REPLACE_URL_TAG     = "##URL##"
	REPLACE_SRI_TAG     = "##SRI_HASH##"
	REPLACE_CROSSORIGIN = "##CROSSORIGIN##"
)

var (
	SRIKinds                    = []string{"sha256", "sha384", "sha512"}
	CROSSORIGIN                 = []string{"anonymous", "use-credentials"}
	defaultInjectionTableHeader = []string{"ID", "Name", "Content", "Created At"}
	url                         string
	sri                         string
	crossorigin                 string
	payload                     string
	protocol                    string
	fields                      []string
	fieldsInjection             []string
	isRaw                       bool
	isReplace                   bool
	useSubdomain                bool
	useHTTPS                    bool
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
			renderInjections(injections)
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
		renderInjections(injections)
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

func renderInjections(injections []api.Injection) {
	rows := buildInjectionsTable(injections)
	if isRaw {
		renderRaw(rows)
		return
	}

	if len(rows) > 0 {
		renderTable(rows)
	} else {
		fmt.Println("No injections found.")
	}
}

func buildInjectionsTable(injections []api.Injection) [][]string {
	var rows [][]string
	if len(fieldsInjection) == 0 {
		fields = defaultInjectionTableHeader
	} else {
		fields = fieldsInjection
	}

	rows = make([][]string, len(injections))

	for i, p := range injections {
		content := ""
		rows[i] = make([]string, 0)
		for _, f := range fields {
			fClean := strings.ReplaceAll(f," ", "")
			fClean = strings.ToLower(fClean)
			switch fClean {
				case "createdat" :
					t := p.CreatedAt.Format("2006-01-02 15:04:05")
					rows[i] = append(rows[i], t)
				case "modifiedat" :
					t := p.ModifiedAt.Format("2006-01-02 15:04:05")
					rows[i] = append(rows[i], t)
				case "content" :
					if isReplace {
						content = p.Content
					} else {
						content = replacePlaceholders(p.Content)
					}
					rows[i] = append(rows[i], content)
				default:
					rows[i] = append(rows[i], getFieldString(p, f))
			}
		}
	}
	return rows
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

	injectionCmd.PersistentFlags().BoolVar(&isReplace, "no-replace", false, "Do not replace placeholder in the injections")
	injectionCmd.PersistentFlags().BoolVar(&useSubdomain, "use-subdomain", true, "Use the subdomain as the payload id (enabled by default)")
	injectionCmd.PersistentFlags().BoolVar(&useHTTPS, "use-https", true, "Use HTTPS (enabled by default)")
	injectionCmd.PersistentFlags().StringVar(&sri, "sri", "sha256", "SRI Type [sha256,sha384,sha512]")
	injectionCmd.PersistentFlags().StringVar(&payload, "payload", "", "Payload ID or Name")
	injectionCmd.PersistentFlags().StringSliceVar(&fieldsInjection, "fields", defaultInjectionTableHeader, "Fields you want to query")
}
