package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	apipkg "github.com/edznux/wonderxss/api"
	httpApi "github.com/edznux/wonderxss/api/http"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage"
	"github.com/edznux/wonderxss/ui"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting web server")

	r := mux.NewRouter()
	api := httpApi.New()
	ui := ui.New()
	api.Routes(r)
	http.Handle("/", r)
	http.HandleFunc("/ui", ui.HandleIndex)

	cfg, err := config.Load("")
	if err != nil {
		log.Fatal(err)
	}
	storage.InitStorage(cfg)

	apipkg.InitApi()

	if cfg.StandaloneHTTPS {
		go func() {
			err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
			if err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
		}()
	}

	go func() {
		err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
