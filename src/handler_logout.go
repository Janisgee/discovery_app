package main

import (
	"net/http"
	"time"
)

func (svr *ApiServer) userLogoutHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get session ID from the cookie
	sessionId, err := r.Cookie("DA_SESSION_ID")
	if err != nil || sessionId == nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	// Invalidate the session (remove from memory and storage)
	delete(svr.memoryUserSessions, sessionId.Value)

	// Clear the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "DA_SESSION_ID",
		Value:    "",
		Expires:  time.Unix(0, 0), // Expired immediately
		HttpOnly: true,
		Path:     "/",
	})

	// Set CORS headers for allow cookies to be sent along with cross-origin requests from server
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Response with a success logout message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}
