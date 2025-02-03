package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func (svr *ApiServer) userBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	type placeRequest struct {
		Username  string `json:"username"`
		PlaceName string `json:"place_name"`
		PlaceID   string `json:"place_id"`
		PlaceText string `json:"place_text"`
		Catagory  string `json:"catagory"`
	}

	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming request body (JSON) into login detail struct
	var newBookmark placeRequest
	err := json.NewDecoder(r.Body).Decode(&newBookmark)
	if err != nil {
		http.Error(w, "Failed to decode user bookmarked place.", http.StatusBadRequest)
		return
	}

	// Get user detail
	user_id := GetCurrentUserId(r)

	// Check if bookmark place has been stored in database
	_, err = svr.bookmarkPlaceService.GetBookmarkPlaceDetails(newBookmark.PlaceID)
	if err != nil {
		// No place found from place id
		// Store place into database

		// Get place details from ChatGPT
		gptResponse, err := svr.locationSvc.GetPlaceDetails(newBookmark.PlaceName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error processing search: %s", err), http.StatusInternalServerError)
			slog.Error("Error in GetPlaceDetails", "error", err)
			return
		}

		err = svr.bookmarkPlaceService.CreatePlaceData(newBookmark.PlaceID, newBookmark.PlaceName, gptResponse.Country, gptResponse.City, newBookmark.Catagory, gptResponse)
		if err != nil {
			http.Error(w, "Failed to create place into place database.", http.StatusBadRequest)
			slog.Error("failed to create place into place database", "error", err)
			return
		}
	}

	// Found place from database.

	// Store bookmark in user_bookmark database
	userBookmarkDetail, err := svr.bookmarkPlaceService.CreateUserBookmark(*user_id, newBookmark.Username, newBookmark.PlaceID, newBookmark.PlaceName, newBookmark.PlaceText)
	if err != nil {
		http.Error(w, "Failed to create user bookmark database for the place.", http.StatusBadRequest)
		slog.Error("failed to create user bookmark database for the place", "error", err)
		return
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"message": newBookmark.PlaceName + " has been bookmarked",
		"UserID":  userBookmarkDetail.UserID,
		"PlaceID": userBookmarkDetail.PlaceID,
	}

	// Set Content-Type to JSON and send a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	// Send JSON response back to client
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}
