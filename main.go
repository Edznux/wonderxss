package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	apipkg "github.com/edznux/wonderxss/api"
	httpApi "github.com/edznux/wonderxss/api/http"
	"github.com/edznux/wonderxss/api/websocket"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/notification"
	"github.com/edznux/wonderxss/storage"
	"github.com/edznux/wonderxss/ui"
	"github.com/gorilla/mux"
)

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

	fmt.Println("WHAT THE FUCK", api.UrlPrefix)
	apiRouter := router.PathPrefix(api.UrlPrefix).Subrouter()
	api.Routes(apiRouter)

	router.HandleFunc("/ws", ws.Handle)
	router.PathPrefix("/").HandlerFunc(ui.HandleIndex)

	apipkg.InitApi()

	if cfg.StandaloneHTTPS {
		go func() {
			fmt.Println("Listenning HTTPS on port :", cfg.HTTPSPOrt)
			err := http.ListenAndServeTLS(":"+strconv.Itoa(cfg.HTTPSPOrt), "server.crt", "server.key", router)
			if err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
		}()
	}

	go func() {
		fmt.Println("Listenning HTTP on port :", cfg.HTTPPOrt)
		err := http.ListenAndServe(":"+strconv.Itoa(cfg.HTTPPOrt), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
		}))

		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	gracefulShutdown()
}

func gracefulShutdown() {
	// handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("shutting down")
	os.Exit(0)
}
