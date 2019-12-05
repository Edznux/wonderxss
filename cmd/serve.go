package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/spf13/cobra"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/api/websocket"
	"github.com/edznux/wonderxss/notification"
	"github.com/edznux/wonderxss/ui"
	"github.com/edznux/wonderxss/webserver"
	"github.com/gorilla/mux"

	httpApi "github.com/edznux/wonderxss/api/http"
)

func gracefulShutdown() {
	// handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("shutting down")
	os.Exit(0)
}

func entrypoint() {
	fmt.Println("Starting web server")

	notification.Setup()
	api.Init()

	router := mux.NewRouter()
	api := httpApi.New()
	ui := ui.New()
	ws := websocket.New()

	apiRouter := router.PathPrefix(api.UrlPrefix).Subrouter()
	api.Routes(apiRouter)

	// Return real payload
	router.HandleFunc("/p/{id}", webserver.HandlePayloadByID)
	router.HandleFunc("/ws", ws.Handle)
	router.HandleFunc("/login", webserver.Login)
	router.HandleFunc("/logout", webserver.Logout)

	router.PathPrefix("/").HandlerFunc(ui.HandleIndex)

	webserver.Serve(router)
	gracefulShutdown()
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		entrypoint()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
