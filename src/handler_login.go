package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// Create a response struct to send back as JSON
	response := map[string]interface{}{
		"message": "Received user login info.",
		"email":   newLogin.Email,
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

	// Hash user password
	hashedPassword, err := hashPassword(newLogin.Password)
	if err != nil {
		fmt.Println("Error in hashing user login password:", err)
	} else {
		fmt.Println("Hashed Password:", hashedPassword)
	}

}

// User creation handler
// func (svr *ApiServer) createUserHandler(w http.ResponseWriter, r *http.Request) {

// 	// Only allow POST requests
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Parse incoming request body (JSON) into User struct
// 	var newUser User
// 	err := json.NewDecoder(r.Body).Decode(&newUser)
// 	if err != nil {
// 		http.Error(w, "Failed to decode user data", http.StatusBadRequest)
// 		return
// 	}

// }
