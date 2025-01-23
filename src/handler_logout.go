package main

import (
	"net/http"
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
	clearSectionCookie(w)

	// Response with a success logout message
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Logged out successfully"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
