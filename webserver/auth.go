package webserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/edznux/wonderxss/crypto"

	"github.com/edznux/wonderxss/api"
)

func Login(w http.ResponseWriter, req *http.Request) {
	log.Printf("Login request")

	res := api.Response{}
	loginParam := req.FormValue("login")
	passwordParam := req.FormValue("password")
	user, err := api.VerifyUserPassword(loginParam, passwordParam)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}

	// Get a new JWT Token if the user is validated
	token, err := crypto.GetJWTToken(user)
	if err != nil {
		log.Println(err)
		res.Error = "Error getting a new token"
		json.NewEncoder(w).Encode(&res)
		return
	}

	res.Data = token
	json.NewEncoder(w).Encode(&res)
}

// Doing logout, server side with JWT is a bit tricky.
// We might want to add blacklisting but it's overkill for this usage IMO
// There should be an enforcement in the validity duration of each token tho.
func Logout(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL : %s, not implemented", r.RequestURI)
	res := api.Response{}
	res.Error = "Not implemented yet"
	json.NewEncoder(w).Encode(&res)
}
