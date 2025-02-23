package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]string{"error": message}
	// Send JSON response back to client
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		// Handle error when encoding response
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func setSectionCookie(w http.ResponseWriter, token string, expiryTime time.Time) {
	// Set the session id cookie in response, not visible to Javascript (HttpOnly)
	http.SetCookie(w, &http.Cookie{
		Name:     "DA_SESSION_ID",
		Value:    token,
		Expires:  expiryTime,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
	})
}

// Clear the session cookie
func clearSectionCookie(w http.ResponseWriter) {
	// Set the session id cookie in response, not visible to Javascript (HttpOnly)
	http.SetCookie(w, &http.Cookie{
		Name:     "DA_SESSION_ID",
		Value:    "",
		Expires:  time.Unix(0, 0), // Expired immediately
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
}
