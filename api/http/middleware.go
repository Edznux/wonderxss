package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	apipkg "github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/crypto"
)

func (api *HTTPApi) jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request URL : %s", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if strings.HasPrefix(r.RequestURI, "/api/v1") {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
func (api *HTTPApi) CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request URL : %s", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if strings.HasPrefix(r.RequestURI, "/api/v1") {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		next.ServeHTTP(w, r)
	})
}

func (api *HTTPApi) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res apipkg.Response

		tokenHeader := r.Header.Get("Authorization")
		if len(tokenHeader) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			res.Error = "Missing Authorization Header"
			json.NewEncoder(w).Encode(&res)
			return
		}

		token := strings.Split(tokenHeader, "Bearer ")[1]
		claims, err := crypto.VerifyJWTToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			res.Error = "Error verifying JWT token: " + err.Error()
			json.NewEncoder(w).Encode(&res)
			return
		}
		fmt.Println(claims)

		next.ServeHTTP(w, r)
	})
}
