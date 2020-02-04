package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/api/http/client"
	"github.com/edznux/wonderxss/api/local"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage"
	"github.com/spf13/cobra"
)

var (
	db         storage.Storage
	remote     bool
	verbose    bool
	currentAPI api.API
	cfg        config.Client
)

func globalFlagsSetup(cmd *cobra.Command, args []string) {
	var err error
	if verbose {
		log.Infoln("Enabling verbose")
		log.SetLevel(log.DebugLevel)
	}

	if remote {
		log.Debugln("Using remote API!")
		cfg, err = config.ReadClientConfig()
		if err != nil {
			log.Fatalln("Coulnd read client config:", err.Error())
		}
		currentAPI = client.New(cfg)
	} else {
		log.Debugln("Using local API")
		currentAPI = local.New()
	}
}

func renderRaw(rows [][]string) {
	for _, row := range rows {
		for _, col := range row {
			fmt.Printf("%s", col)
		}
		fmt.Printf("\n")
	}
}

func renderTable(rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(fields)
	table.AppendBulk(rows)
	table.Render()
}

func getFieldString(i interface{}, field string) string {
	r := reflect.ValueOf(i)
	f := reflect.Indirect(r).FieldByNameFunc(func(n string) bool { return strings.ToLower(n) == strings.ToLower(field) })
	return f.String()
}

var rootCmd = &cobra.Command{
	Use:   "wonderxss",
	Short: "WonderXSS is a pentest tool for discovering Blind XSSs",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		globalFlagsSetup(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		// No arguments called. Abort and print help
		cmd.Help()
	},
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Get the health of the application",
	Run: func(cmd *cobra.Command, args []string) {
		// No arguments called. Abort and print help
		result, err := currentAPI.GetHealth()
		if err != nil {
			fmt.Println("Error :", err.Error())
		}
		fmt.Println(result)
	},
}

//Execute is the "main" function for the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&remote, "remote", "r", false, "Use remote API instead of local")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose (debug) logs")
	rootCmd.PersistentFlags().BoolVar(&isRaw, "raw", false, "Output without any formating (good for scripting)")
	rootCmd.AddCommand(healthCmd)
}
