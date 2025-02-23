package api

import (
	"discoveryweb/util"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

func (svr *ApiServer) userResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	type resetPw struct {
		NewPw       string `json:"newPw"`
		ConfirmPw   string `json:"confirmPw"`
		PwResetCode string `json:"pwResetCode"`
	}
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Parse incoming request body (JSON) into login detail struct
	var resetInfo resetPw
	err := json.NewDecoder(r.Body).Decode(&resetInfo)
	if err != nil {

		http.Error(w, "Failed to decode user reset password data", http.StatusBadRequest)
		return
	}

	// Check password strength
	err = util.CheckPasswordStrength(resetInfo.NewPw)
	if err != nil {
		errMessage := fmt.Sprintf("%v", err)
		sendErrorResponse(w, http.StatusBadRequest, errMessage)
	}

	// Trim space for new password
	trimedPassword := strings.TrimSpace(resetInfo.NewPw)

	// Validate connection link if it has record in database
	// Get user email from pw reset code if pw code is still valid.
	pwCode := strings.TrimSuffix(resetInfo.PwResetCode, "/")

	userEmail, err := svr.userSvc.GetUserEmailFromEmailPw(pwCode)
	if err != nil {
		svr.UnhandledError(err)
		if err.Error() == "sql: no rows in result set" {
			http.Error(w, "Unauthorized reset link", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Failed to get user email from pw reset code", http.StatusInternalServerError)
		return
	}

	// Get user info from user email
	userInfo, err := svr.userSvc.ResetUserPw(userEmail, trimedPassword, pwCode)
	if err != nil {
		slog.Warn("Fail to get user info from reset user password", "error", err)
		http.Error(w, "Fail to get user info from reset user password", http.StatusUnauthorized)
		return
	}

	// Create a response struct to send back as JSON
	response := map[string]interface{}{
		"message":  "Updated user new password successfully.",
		"username": userInfo.Username,
		"email":    userInfo.Email,
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
