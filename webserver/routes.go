package webserver

import "github.com/gorilla/mux"

// Routes defines the non-api routes (User interface and Payload serving using path)
func (ui *UI) Routes(router *mux.Router) {
	router.HandleFunc("/p/{id}", ui.HandlePayloadByID)
	router.HandleFunc("/login", ui.Login)
	router.HandleFunc("/logout", ui.Logout)
	router.HandleFunc("/otp/new", ui.GenerateOTPSecret).Methods("GET")
	router.HandleFunc("/otp/new", ui.RegisterOTP).Methods("POST")

	router.PathPrefix("/").HandlerFunc(ui.HandleIndex)
}
