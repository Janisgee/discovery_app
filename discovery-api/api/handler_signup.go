package api

import (
	"discoveryweb/service/user"
	"discoveryweb/util"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	passwordvalidator "github.com/wagslane/go-password-validator"
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

	// Check password strength
	const minEntropyBits = 60
	err = passwordvalidator.Validate(newSignup.Password, minEntropyBits)
	if err != nil {
		slog.Warn("Password is not strong enough.Plesae signup again with a stronger password. Suggestion:", "error", err)
		errMessage := fmt.Sprintf("%v", err)
		sendErrorResponse(w, http.StatusBadRequest, errMessage)
		return
	}
	// Check valid email structure
	splitEmail := strings.Split(newSignup.Email, "@")
	if len(splitEmail) < 2 {
		errMessage := "please fill in a valid email for signing up."
		sendErrorResponse(w, http.StatusBadRequest, errMessage)
		return
	}

	// Check password strength
	err = util.CheckPasswordStrength(newSignup.Password)
	if err != nil {
		errMessage := fmt.Sprintf("%v", err)
		sendErrorResponse(w, http.StatusBadRequest, errMessage)
	}

	// Trim space and store lowercase from input
	trimedUsername := strings.TrimSpace(newSignup.Username)
	trimedEmail := strings.TrimSpace(strings.ToLower(newSignup.Email))
	trimedPassword := strings.TrimSpace(newSignup.Password)

	// Call CreateUser from the user service to create the user
	_, err = svr.userSvc.CreateUser(trimedUsername, trimedEmail, trimedPassword)
	if err != nil {
		if errors.Is(err, user.ErrEmailInUse) {
			//Do Nothing
		} else {
			svr.UnhandledError(err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}
	}

	// Create a response struct to send back as JSON
	response := map[string]interface{}{
		"message": "Received user signup info.",
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
