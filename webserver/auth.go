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
	Token string `json:"token"`
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

// ValidateOTP is handling the login attempt with 2FA.
// It parses a new OTP Token (6 number digit) and validates it against one we generates
func ValidateOTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("2FA Request")
	otpToken := OtpToken{}
	res := api.Response{}

	// Get the OTP Token (6 digits)
	err := json.NewDecoder(req.Body).Decode(&otpToken)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	// Get the JWT token so we can extract the user id etc...
	jwtToken, err := getBearer(w, req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	// Verify the JWT Token first, if it's not valid reject the request
	decodedToken, err := crypto.VerifyJWTToken(jwtToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		res.Error = "Error verifying JWT token: " + err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}

	jwtToken2FA, err := verifyOTP(decodedToken, otpToken.Token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		res.Error = "Error verifying JWT token: " + err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	// Get the validated new JWT Token and send it back
	res.Data = jwtToken2FA
	json.NewEncoder(w).Encode(&res)
}

// RegisterOTP recieve a new shared secret for TOTP and saves it for a given user
func RegisterOTP(w http.ResponseWriter, req *http.Request) {

	otpToken := OtpToken{}
	res := api.Response{}

	err := json.NewDecoder(req.Body).Decode(&otpToken)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	bearer, err := getBearer(w, req)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	claims, err := crypto.VerifyJWTToken(bearer)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}

	decodedTokenMap := claims.(jwt.MapClaims)
	userID := decodedTokenMap["user_id"].(string)

	user, err := api.CreateOTP(userID, otpToken.Token)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	if !user.TwoFactorEnabled {
		res.Error = "2FA wasn't saved properly"
		json.NewEncoder(w).Encode(&res)
		return
	}
	res.Data = "OK"
	json.NewEncoder(w).Encode(&res)
}

// GenerateOTPSecret generates a new secrets (80 bit base32 encoded)
// FIXME: Shouldn't this only be done client side ?
// Maybe dont trust client side crypto ? idk
func GenerateOTPSecret(w http.ResponseWriter, req *http.Request) {
	token, err := crypto.GenerateOTPSecret()

	res := api.Response{}
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(&res)
		return
	}
	res.Data = OtpToken{Token: token}
	json.NewEncoder(w).Encode(&res)
}

// Login is the http handler function for user login
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

	// Get a new JWT Token if the user is validated.
	// The second parameter MUST be false since we didn't checked the 2FA yet.
	token, err := crypto.GetJWTToken(user, false)
	if err != nil {
		log.Println(err)
		res.Error = "Error getting a new token"
		json.NewEncoder(w).Encode(&res)
		return
	}

	res.Data = token
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

// verifyOTP takes the bearerToken and the OTPToken.
// It needs the bearer token to ensure the user id is valid and has the 2FA enabled.
// With this valid bearer, it will also get the User and it's TOTPSecret stored in DB
// It will returns a new JWTToken with the "2FAVerified" field set to true (if successful)
func verifyOTP(decodedToken jwt.Claims, token string) (string, error) {

	decodedTokenMap := decodedToken.(jwt.MapClaims)

	// Get the user first, so we can double check it has 2FA enabled, it exists etc...
	userID := decodedTokenMap["user_id"].(string)
	user, err := api.GetUser(userID)
	if err != nil {
		return "", err
	}

	secret := user.TOTPSecret
	if len(secret) == 0 {
		return "", fmt.Errorf("Empty TOTPSecret")
	}

	otpc := &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  3,
		HotpCounter: 0,
	}

	decodedTokenMap["2FAVerified"], err = otpc.Authenticate(token)
	if err != nil {
		return "", err
	}

	if decodedTokenMap["2FAVerified"] == false {
		return "", fmt.Errorf("Invalid one-time password")
	}

	jwToken, _ := crypto.GetJWTToken(user, true)
	return jwToken, nil
}
