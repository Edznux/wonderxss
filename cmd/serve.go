package cmd

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/edznux/wonderxss/api/websocket"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/notification"
	"github.com/edznux/wonderxss/webserver"
	"github.com/gorilla/mux"

	httpApi "github.com/edznux/wonderxss/api/http/server"
)

var (
	httpPort  int16 = 80
	httpsPort int16 = 443
)

func gracefulShutdown() {
	// handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Info("shutting down")
	os.Exit(0)
}

func entrypoint() {
	config.Current.HTTPPOrt	= int(httpPort)
	config.Current.HTTPSPOrt= int(httpsPort)
	notification.Setup()

	router := mux.NewRouter()
	api := httpApi.New()
	ui := webserver.New()
	ws := websocket.New()

	router.Use(ui.LoggingMiddleware)
	apiRouter := router.PathPrefix(api.UrlPrefix).Subrouter()
	api.Routes(apiRouter)
	// Return real payload
	router.HandleFunc("/ws", ws.Handle)
	ui.Serve(router)
	gracefulShutdown()
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"server"},
	Short:   "Start the application server",
	Run: func(cmd *cobra.Command, args []string) {
		entrypoint()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().Int16Var(&httpPort, "http-port", 80, "HTTP Port")
	serveCmd.PersistentFlags().Int16Var(&httpsPort, "https-port", 443, "HTTPS Port")
}
