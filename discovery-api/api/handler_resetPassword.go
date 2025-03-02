package api

import (
	"discoveryweb/service/user"
	"discoveryweb/util"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
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
		http.Error(w, "password is not secure enough", http.StatusBadRequest)
		return
	}

	if resetInfo.NewPw != resetInfo.ConfirmPw {
		http.Error(w, "password and confirm do not match", http.StatusBadRequest)
		return
	}

	if resetInfo.NewPw != strings.TrimSpace(resetInfo.NewPw) {
		http.Error(w, "password cannot start or end with spaces", http.StatusBadRequest)
		return
	}

	resetKey, err := uuid.Parse(resetInfo.PwResetCode)
	if err != nil {
		http.Error(w, "invalid or expired reset code", http.StatusUnauthorized)
		return
	}

	err = svr.userSvc.CompleteUserPasswordReset(resetKey, resetInfo.NewPw)
	if err != nil {
		if errors.Is(err, user.ErrPasswordResetInvalid) {
			http.Error(w, "invalid or expired reset code", http.StatusUnauthorized)
		} else {
			http.Error(w, "unexpected error has occurred", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK) // 200 OK
}
