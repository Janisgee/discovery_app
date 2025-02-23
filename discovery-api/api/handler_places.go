package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (svr *ApiServer) autocompleteCitiesSearch(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	cityQuery := r.URL.Query().Get("search")
	locale := r.URL.Query().Get("locale")
	if locale == "" {
		locale = "en_US"
	}

	if cityQuery == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	cityOptions, err := svr.placesService.AutocompleteCities(cityQuery, locale)

	if err != nil {
		slog.Error("Error completing city search", "err", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// Create the JSON response
	jsData, err := json.Marshal(cityOptions)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsData)
	if err != nil {
		slog.Warn("Failed to write JS to http response", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
