package api

import (
	"discoveryweb/service/user"
	"encoding/json"
	"errors"
	"net/http"
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

	err = svr.userSvc.BeginUserPasswordReset(request.Email)
	if err != nil {
		if errors.Is(err, user.ErrNoUser) {
			// Do nothing
		} else {
			svr.UnhandledError(err)
			http.Error(w, "Unexpected error", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK) // 200 OK
}
