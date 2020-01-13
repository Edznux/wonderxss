package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/edznux/wonderxss/config"
	"github.com/gorilla/mux"
)

func (ui *UI) Serve(router *mux.Router) {
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
		err := http.ListenAndServe(":"+strconv.Itoa(cfg.HTTPPOrt), router)

		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}
