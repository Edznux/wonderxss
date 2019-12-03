package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/edznux/wonderxss/api/websocket"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/notification"
	"github.com/edznux/wonderxss/storage"
	"github.com/edznux/wonderxss/ui"
	"github.com/edznux/wonderxss/webserver"
	"github.com/gorilla/mux"

	apipkg "github.com/edznux/wonderxss/api"
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

func main() {
	fmt.Println("Starting web server")
	cfg, err := config.Load("")
	if err != nil {
		log.Fatal(err)
	}
	storage.InitStorage(cfg)

	notification.Setup(cfg)

	router := mux.NewRouter()
	api := httpApi.New()
	ui := ui.New()
	ws := websocket.New()

	apiRouter := router.PathPrefix(api.UrlPrefix).Subrouter()
	api.Routes(apiRouter)

	// Return real payload
	router.HandleFunc("/p/{id}", webserver.HandlePayloadByID)

	router.HandleFunc("/ws", ws.Handle)
	router.PathPrefix("/").HandlerFunc(ui.HandleIndex)

	apipkg.InitApi()

	webserver.Serve(cfg, router)
	gracefulShutdown()
}
