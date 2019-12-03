package webserver

import (
	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	// Authentication
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("POST")
}
