package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/edznux/wonderxss/crypto"

	"github.com/edznux/wonderxss/api"
)

func Login(w http.ResponseWriter, req *http.Request) {
	log.Printf("Login request")
	log.Printf("r: %+v\n", req)

	res := api.Response{}
	loginParam := req.FormValue("login")
	passwordParam := req.FormValue("password")

	fmt.Println("login, passwd", loginParam, passwordParam)
	user, err := api.VerifyUserPassword(loginParam, passwordParam)
	if err != nil {
		res.Code = 0
		res.Message = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}

	// Get a new JWT Token if the user is validated
	token, err := crypto.GetJWTToken(user)
	if err != nil {
		log.Println(err)
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
