package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/edznux/wonderxss/config"
	"github.com/gorilla/mux"
)

func Serve(router *mux.Router) {
	cfg := config.Current
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
}
