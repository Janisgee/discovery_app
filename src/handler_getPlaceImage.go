package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// GetImageURL

func (svr *ApiServer) getPlaceImageURL(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Place string `json:"place_name"`
	}
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Parse incoming request body (JSON) into login detail struct
	var placeInfo request
	err := json.NewDecoder(r.Body).Decode(&placeInfo)
	if err != nil {
		http.Error(w, "Failed to decode place data", http.StatusBadRequest)
		return
	}

	placeData, err := svr.imgSvc.GetImageURL(placeInfo.Place)
	if err != nil {
		slog.Warn("Failed to decode place img url", "error", err)
		http.Error(w, "Failed to decode place img url", http.StatusBadRequest)
		return
	}

	fmt.Println(placeData)
	// Create a response struct to send back as JSON
	response := map[string]interface{}{
		"place_name": placeInfo.Place,
		"image_id":   placeData.ImageID,
		"image_url":  placeData.ImageURL,
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
