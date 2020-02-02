package webserver

import "github.com/gorilla/mux"

func (ui *UI) Routes(router *mux.Router) {
	router.HandleFunc("/p/{id}", ui.HandlePayloadByID)
	router.HandleFunc("/login", ui.Login)
	router.HandleFunc("/logout", ui.Logout)
	router.HandleFunc("/otp/new", ui.GenerateOTPSecret).Methods("GET")
	router.HandleFunc("/otp/new", ui.RegisterOTP).Methods("POST")

	router.PathPrefix("/").HandlerFunc(ui.HandleIndex)
}
