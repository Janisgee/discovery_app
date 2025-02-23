package api

import (
	"net/http"
)

func (svr *ApiServer) userLogoutHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get session ID from the cookie
	sessionId := sessionIdFromCookie(r)

	svr.sessionSvc.DeleteSession(*sessionId)

	// Clear the session cookie
	clearSectionCookie(w)

	// Response with a success logout message
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Logged out successfully"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
