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
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	api := httpApi.New()
	api.Routes(apiRouter)

	ui := ui.New()
	ws := websocket.New()
	http.HandleFunc("/", ui.HandleIndex)
	http.HandleFunc("/ws", ws.Handle)

	apipkg.InitApi()

	if cfg.StandaloneHTTPS {
		go func() {
			fmt.Println("Listenning HTTPS on port :", cfg.HTTPSPOrt)
			err := http.ListenAndServeTLS(":"+strconv.Itoa(cfg.HTTPSPOrt), "server.crt", "server.key", nil)
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
