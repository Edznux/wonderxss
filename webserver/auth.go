package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgryski/dgoogauth"
	"github.com/edznux/wonderxss/crypto"

	"github.com/edznux/wonderxss/api"
)

type OtpToken struct {
	Token  string `json:"token,omitempty"`
	Secret string `json:"secret,omitempty"`
}

func getBearer(w http.ResponseWriter, req *http.Request) (string, error) {
	// Get the header
	tokenHeader := req.Header.Get("Authorization")
	if len(tokenHeader) == 0 {
		return "", fmt.Errorf("Missing Authorization Header")
	}
	bearer := strings.Split(tokenHeader, "Bearer ")
	if len(bearer) != 2 {
		return "", fmt.Errorf("Error verifying JWT token: Invalid token")
	}
	return bearer[1], nil
}

// RegisterOTP recieve a new shared secret for TOTP and saves it for a given user
func RegisterOTP(w http.ResponseWriter, req *http.Request) {

	otpToken := OtpToken{}
	res := api.Response{}

	err := json.NewDecoder(req.Body).Decode(&otpToken)
	if err != nil {
		log.Println("Error while decoding OtpToken:", err)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	bearer, err := getBearer(w, req)
	if err != nil {
		log.Println("getBearer error:", err)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	claims, err := crypto.VerifyJWTToken(bearer)
	if err != nil {
		log.Println("Error verifying the JWT Token:", err)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}

	decodedTokenMap := claims.(jwt.MapClaims)
	userID := decodedTokenMap["user_id"].(string)

	user, err := api.CreateOTP(userID, otpToken.Secret)
	if err != nil {
		log.Println(err)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	if !user.TwoFactorEnabled {
		res.Error = "2FA wasn't saved properly"
		log.Println(res.Error)
		json.NewEncoder(w).Encode(&res)
		return
	}
	// Verify the token with the secret
	jwtTokenVerified, err := verifyOTP(user.TOTPSecret, otpToken.Token)
	if err != nil {
		log.Println("Could not verify the newly registered OTP:", err)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	res.Data = jwtTokenVerified
	json.NewEncoder(w).Encode(&res)
}

// GenerateOTPSecret generates a new secrets (80 bit base32 encoded)
// FIXME: Shouldn't this only be done client side ?
// Maybe dont trust client side crypto ? idk
func GenerateOTPSecret(w http.ResponseWriter, req *http.Request) {
	token, err := crypto.GenerateOTPSecret()

	res := api.Response{}
	if err != nil {
		log.Println(err)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	res.Data = OtpToken{Secret: token}
	json.NewEncoder(w).Encode(&res)
}

// Login is the http handler function for user login
func Login(w http.ResponseWriter, req *http.Request) {
	log.Printf("Login request")
	var OTPOk bool

	res := api.Response{}
	loginParam := req.FormValue("login")
	passwordParam := req.FormValue("password")

	user, err := api.VerifyUserPassword(loginParam, passwordParam)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	if user.TwoFactorEnabled {
		OTPToken := req.FormValue("token")
		OTPOk, err = verifyOTP(user.TOTPSecret, OTPToken)
		if err != nil {
			res.Error = err.Error()
			json.NewEncoder(w).Encode(&res)
			return
		}
		if !OTPOk {
			res.Error = "Invalid OTP"
			json.NewEncoder(w).Encode(&res)
			return
		}
	}
	// Get a new JWT Token if the user is validated.
	jwtToken, err := crypto.GetJWTToken(user)
	if err != nil {
		log.Println(err)
		res.Error = "Error getting a new token"
		json.NewEncoder(w).Encode(&res)
		return
	}

	res.Data = jwtToken
	json.NewEncoder(w).Encode(&res)
}

// Logout is the http handler function.
// Login out with JWT is a bit tricky since there is no real way of invalidating a JWT.
// We might want to add blacklisting but it's overkill for this usage IMO
// There MUST be an enforcement in the validity duration of each token tho.
func Logout(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL : %s, not implemented", r.RequestURI)
	res := api.Response{}
	res.Error = "Not implemented yet"
	json.NewEncoder(w).Encode(&res)
}

// verifyOTP takes the user's OTP secret and the OTPToken. It will return true if it's valid.
func verifyOTP(secret string, token string) (bool, error) {

	if len(secret) == 0 {
		return false, fmt.Errorf("Empty TOTPSecret")
	}

	otpc := &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  3,
		HotpCounter: 0,
	}

	verified, err := otpc.Authenticate(token)
	if err != nil {
		log.Println("VerifyOTP failed authenticate:", err)
		return false, err
	}

	if !verified {
		return false, fmt.Errorf("Invalid one-time password")
	}

	return true, nil
}
