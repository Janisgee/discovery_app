package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

type SearchPlaceDetails struct {
	Place string `json:"place"`
}

// handleSearchPlaceDetails processes the incoming POST request for a place search

func (svr *ApiServer) handleSearchPlaceDetails(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON data
	var input SearchPlaceDetails

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Process the input data
	fmt.Printf("Received place name (line32): %s \n", input.Place)

	// Handle the page and get place details for the provided place

	response, err := svr.locationSvc.GetPlaceDetails(input.Place)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing search: %s", err), http.StatusInternalServerError)
		slog.Error("Error in GetPlaceDetails", "error", err)
		return
	}

	// Create the JSON response
	jsData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	_, err = w.Write(jsData)
	if err != nil {
		log.Printf("Failed to write JSON response: %s\n", err)
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
	}
}
