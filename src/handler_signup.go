package main

import (
	"encoding/json"
	"net/http"
)

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

	// Validate the input fields (basic checks)
	if newSignup.Username == "" || newSignup.Email == "" || newSignup.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Call CreateUser from the user service to create the user
	user, err := svr.userSvc.CreateUser(newSignup.Username, newSignup.Email, newSignup.Password)
	if err != nil {
		svr.UnhandledError(err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Create a response struct to send back as JSON
	response := map[string]interface{}{
		"message":  "Received user signup info.",
		"username": user.Username,
		"email":    user.Email,
	}

	// Set Content-Type to JSON and send a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	// Send JSON response back to client
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// Handle error when encoding response
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}

}
