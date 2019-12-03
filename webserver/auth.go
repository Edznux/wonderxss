package webserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/edznux/wonderxss/crypto"

	"github.com/edznux/wonderxss/api"
	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL : %s, not implemented", r.RequestURI)
	res := api.Response{}
	vars := mux.Vars(r)
	loginParam := vars["login"]
	passwordParam := vars["password"]
	
	user, err := api.VerifyUserPassword(loginParam, passwordParam)
	if err != nil {
		res.Code = 0
		res.Message = err
		json.NewEncoder(w).Encode(&res)
		return 
	}

	// Get a new JWT Token if the user is validated
	token, err := crypto.GetJWTToken(user)
	if err != nil {
		res.Code = 0
		res.Message = "Error getting a new token"
		json.NewEncoder(w).Encode(&res)
		return
	}

	res.Code = 0
	res.Message = "OK"
	res.Data = token
	json.NewEncoder(w).Encode(&res)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL : %s, not implemented", r.RequestURI)
	res := api.Response{}
	res.Code = 0
	res.Message = "Not implemented yet"
	json.NewEncoder(w).Encode(&res)
}
