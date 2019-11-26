package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/edznux/wonderxss/api"
	apipkg "github.com/edznux/wonderxss/api"
	httpApi "github.com/edznux/wonderxss/api/http"
	"github.com/edznux/wonderxss/api/websocket"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/notification"
	"github.com/edznux/wonderxss/storage"
	"github.com/edznux/wonderxss/ui"
	"github.com/gorilla/mux"
)

// This shouldn't be here. I know. But it doesn't belong to api/http either
// It's not using the /api/v1 prefix. (And we don't want that because need short payload)
// But putting this alone in a /route folder or even /http feels a bit wierd too.
func handlePayloadByID(w http.ResponseWriter, req *http.Request) {
	var err error

	params := mux.Vars(req)
	id := params["id"]
	text, err := api.ServePayload(id)
	if err != nil {
		fmt.Printf("Could not get payload to be served as a /p/%s, error : %s\n", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(text))
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
	router.HandleFunc("/p/{id}", handlePayloadByID)

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
