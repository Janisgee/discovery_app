package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (svr *ApiServer) userForgetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	type requestEmail struct {
		Email string `json:"email"`
	}
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming request body (JSON) into login detail struct
	var request requestEmail
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode user forget password data", http.StatusBadRequest)
		return
	}

	// Validate the input fields (basic checks)
	if request.Email == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Check if user provided email valid
	userInfo, err := svr.userSvc.GetUserInfo(request.Email)
	if err != nil {
		slog.Warn("Fail to reset user password. Please try again later", "error", err)
		http.Error(w, "Fail to reset user password. Please try again later", http.StatusUnauthorized)
		return
	}

	// Generate random email password
	alphaNumRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	emailVerRandRune := make([]rune, 64)

	for i := 0; i < 64; i++ {
		/* #nosec */
		emailVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}
	fmt.Println("change password emailVerRandRune:", emailVerRandRune)
	emailVerPassword := string(emailVerRandRune)

	// Hash generated email password for db

	emailVerPWhash, err := bcrypt.GenerateFromPassword([]byte(emailVerPassword), bcrypt.DefaultCost)
	if err != nil {
		slog.Warn("bcrypt err:", "error", err)
		http.Error(w, "Fail to generate hash email password", http.StatusUnauthorized)
		return
	}

	// Store hashedEmailPw into UserEmailPw table
	err = svr.userSvc.CreateUserEmailPw(userInfo.Email, string(emailVerPWhash), userInfo.ID)

	if err != nil {
		slog.Warn("Fail to create user email password", "error", err)
		http.Error(w, "Fail to create user email password", http.StatusUnauthorized)
		return
	}

	//Structure the retrieve password link
	retrievePwLink := "http://mysite.com/forgetPwChange?" + "evpw=" + string(emailVerPWhash) + "/"

	// Send reset account email to user
	err = svr.emailSvc.SendPasswordResetEmail(userInfo.Username, userInfo.Email, retrievePwLink)
	if err != nil {
		slog.Warn("Fail to send reset account email to user", "error", err)
		http.Error(w, "Fail to send reset account email to user", http.StatusUnauthorized)
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"email":               userInfo.Email,
		"hashedEmailPassword": retrievePwLink,
	}
	// Set Content-Type to JSON and send a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	// Send JSON response back to client
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}

}
