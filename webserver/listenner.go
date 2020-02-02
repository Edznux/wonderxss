package webserver

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/config"
	"github.com/gorilla/mux"
)

func (ui *UI) Serve(router *mux.Router) {
	ui.Routes(router)

	cfg := config.Current
	if cfg.StandaloneHTTPS {
		go func() {
			log.Println("Listening HTTPS on port :", cfg.HTTPSPOrt)
			err := http.ListenAndServeTLS(":"+strconv.Itoa(cfg.HTTPSPOrt), "server.crt", "server.key", router)
			if err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
		}()
	}

	go func() {
		log.Println("Listening HTTP on port :", cfg.HTTPPOrt)
		err := http.ListenAndServe(":"+strconv.Itoa(cfg.HTTPPOrt), router)

		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}
