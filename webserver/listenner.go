package webserver

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/config"
	"github.com/gorilla/mux"
)

//Serve run the webserve on http & https (if enabled)
func (ui *UI) Serve(router *mux.Router) {
	ui.Routes(router)

	cfg := config.Current
	if cfg.StandaloneHTTPS {
		go func() {
			log.Printf("Listening HTTPS on %s:%d\n", cfg.ListeningAddress, cfg.HTTPSPOrt)
			err := http.ListenAndServeTLS(cfg.ListeningAddress+":"+strconv.Itoa(cfg.HTTPSPOrt), "server.crt", "server.key", router)
			if err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
		}()
	}

	go func() {
		log.Printf("Listening HTTP on on %s:%d\n", cfg.ListeningAddress, cfg.HTTPPOrt)
		err := http.ListenAndServe(cfg.ListeningAddress+":"+strconv.Itoa(cfg.HTTPPOrt), router)

		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}
