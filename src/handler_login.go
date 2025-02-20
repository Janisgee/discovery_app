package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (svr *ApiServer) userLoginHandler(w http.ResponseWriter, r *http.Request) {
	type loginDetail struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming request body (JSON) into login detail struct
	var newLogin loginDetail
	err := json.NewDecoder(r.Body).Decode(&newLogin)
	if err != nil {
		http.Error(w, "Failed to decode user login data", http.StatusBadRequest)
		return
	}

	// Validate the input fields (basic checks)
	if newLogin.Email == "" || newLogin.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	validUserId, err := svr.userSvc.VerifyUserLogin(newLogin.Email, newLogin.Password)
	if err != nil {
		slog.Warn("Failed login attempt", "error", err)
		http.Error(w, "Fail to login user", http.StatusUnauthorized)
		return
	}

	// Create a session token
	token := uuid.NewString()
	expiryTime := time.Now().Add(600 * time.Second) // expires in 10 min

	// Store the session in server memory
	svr.memoryUserSessions[token] = UserSession{
		userId:     validUserId,
		expiryTime: expiryTime,
	}

	setSectionCookie(w, token, expiryTime)

	// Get username from user input email
	userInfo, err := svr.userSvc.GetUserInfo(newLogin.Email)
	if err != nil {
		slog.Warn("Fail to get user info from login input email", "error", err)
		http.Error(w, "Fail to get user info from login input email", http.StatusUnauthorized)
		return
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"message":  "User has been logged in.",
		"username": userInfo.Username,
		"email":    userInfo.Email,
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
