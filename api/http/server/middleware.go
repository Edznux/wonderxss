package http

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"strings"

	apipkg "github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/crypto"
)

func (api *HTTPApi) jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Request URL : %s", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if strings.HasPrefix(r.RequestURI, "/api/v1") {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

func (api *HTTPApi) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Request URL : %s", r.RequestURI)
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
		bearer := strings.Split(tokenHeader, "Bearer ")
		if len(bearer) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			res.Error = "Error verifying JWT token: Invalid token"
			json.NewEncoder(w).Encode(&res)
			return
		}
		token := bearer[1]

		claims, err := crypto.VerifyJWTToken(token, config.Current.JWTToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			res.Error = "Error verifying JWT token: " + err.Error()
			json.NewEncoder(w).Encode(&res)
			return
		}
		log.Debugf("Claims: %+v", claims)

		next.ServeHTTP(w, r)
	})
}
