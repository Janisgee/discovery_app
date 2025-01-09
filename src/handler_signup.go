package main

import (
	"encoding/json"
	"net/http"
)

// User struct to hold input data
type User struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Hashed_password string `json:"password"`
}

func (svr *ApiServer) userSignupHandler(w http.ResponseWriter, r *http.Request) {
	type signupDetail struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming request body (JSON) into login detail struct
	var newSignup signupDetail
	err := json.NewDecoder(r.Body).Decode(&newSignup)
	if err != nil {
		http.Error(w, "Failed to decode user signup data", http.StatusBadRequest)
		return
	}

	// Create a response struct to send back as JSON
	response := map[string]interface{}{
		"message":  "Received user signup info.",
		"username": newSignup.Username,
		"email":    newSignup.Email,
	}

	// Set Content-Type to JSON and send a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	// Send JSON response
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// Handle error when encoding response
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}

}
