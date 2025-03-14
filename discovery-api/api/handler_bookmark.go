package api

import (
	"discoveryweb/internal/database"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func (svr *ApiServer) userUnBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	type unBookmark struct {
		PlaceID   string `json:"place_id"`
		PlaceName string `json:"place_name"`
	}

	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming request body (JSON) into login detail struct
	var unBookmarkFromUser unBookmark
	err := json.NewDecoder(r.Body).Decode(&unBookmarkFromUser)
	if err != nil {
		http.Error(w, "Failed to decode user bookmarked place.", http.StatusBadRequest)
		return
	}

	// Get user detail
	user_id := GetCurrentUserId(r)

	// Delete bookmark by user id and place id
	deletedBookmark, err := svr.bookmarkPlaceService.DeleteUserBookmark(*user_id, unBookmarkFromUser.PlaceID)
	if err != nil {
		slog.Warn("Fail to delete user bookmark place", "error", err)
		http.Error(w, "Fail to delete user bookmark place.", http.StatusBadRequest)
		return
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"message": unBookmarkFromUser.PlaceName + " has been unbookmarked",
		"UserID":  deletedBookmark.UserID,
		"PlaceID": deletedBookmark.PlaceID,
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

func (svr *ApiServer) userGetAllBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get user detail
	user_id := GetCurrentUserId(r)

	// Get all bookmark list from user
	placeIDList, err := svr.bookmarkPlaceService.GetAllBookmarkedPlace(*user_id)
	if err != nil {
		http.Error(w, "Fail to get all place ID from user bookmark database", http.StatusMethodNotAllowed)
		return
	}
	// If no rows found, return an empty slice
	if len(placeIDList) == 0 {
		// Optionally log that no bookmarks were found
		slog.Info("No bookmarked places found for user", "user_id", user_id)
		placeIDList = []database.GetAllUserBookmarkPlaceIDRow{}
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"message":         "Place ID that user has been bookmarked",
		"BookmarkedPlace": placeIDList,
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

func (svr *ApiServer) userGetAllBookmarkByCityHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	type request struct {
		City string `json:"city"`
	}
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Parse incoming request body (JSON) into login detail struct
	var locationInfo request
	err := json.NewDecoder(r.Body).Decode(&locationInfo)
	if err != nil {
		http.Error(w, "Failed to decode user bookmark place name.", http.StatusBadRequest)
		return
	}

	// Get user detail
	user_id := GetCurrentUserId(r)

	// Get all bookmark list from user
	placeList, err := svr.bookmarkPlaceService.GetAllBookmarkedCity(*user_id, locationInfo.City)
	if err != nil {
		http.Error(w, "Fail to get all city information from user bookmark database", http.StatusMethodNotAllowed)
		return
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"PlaceList": placeList,
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

func (svr *ApiServer) userBookmarkByPlaceNameHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Username  string `json:"username"`
		PlaceName string `json:"place_name"`
		PlaceID   string `json:"place_id"`
	}
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Parse incoming request body (JSON) into login detail struct
	var bookmarkRequest request
	err := json.NewDecoder(r.Body).Decode(&bookmarkRequest)
	if err != nil {
		http.Error(w, "Failed to decode user bookmark place name.", http.StatusBadRequest)
		return
	}

	// Get place information from db
	placeInfo, err := svr.bookmarkPlaceService.GetPlaceDatabaseDetails(bookmarkRequest.PlaceID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting place details from database: %s", err), http.StatusInternalServerError)
		slog.Error("Error getting place details from database", "error", err)
		return
	}

	// Get user detail
	user_id := GetCurrentUserId(r)

	// Check if place has been bookmarked by user (Return true or false)
	result, _ := svr.bookmarkPlaceService.CheckPlaceHasBookmarkedByUser(bookmarkRequest.PlaceID, *user_id)
	if !result {

		// Bookmark place
		bookmarkPlace, err := svr.bookmarkPlaceService.CreateUserBookmark(*user_id, bookmarkRequest.Username, bookmarkRequest.PlaceID, bookmarkRequest.PlaceName, placeInfo.Category, placeInfo.PlaceDetail.Description)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating user bookmark at place detail page: %s", err), http.StatusInternalServerError)
			slog.Error("Error creating user bookmark at place detail page", "error", err)
			return
		}

		// Console response struct to send back as JSON
		response := map[string]interface{}{
			"message": bookmarkRequest.PlaceName + " has been bookmarked",
			"UserID":  bookmarkPlace.UserID,
			"PlaceID": bookmarkPlace.PlaceID,
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
		return
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"message": placeInfo.PlaceName + " has already been bookmarked before",
		"UserID":  user_id,
		"PlaceID": bookmarkRequest.PlaceID,
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

	// Check if place has been bookmarked by user (Return true or false)
	result, _ := svr.bookmarkPlaceService.CheckPlaceHasBookmarkedByUser(newBookmark.PlaceID, *user_id)
	if !result {
		//No bookmark found from user

		// Check if place has been stored in database
		_, err = svr.bookmarkPlaceService.GetPlaceDatabaseDetails(newBookmark.PlaceID)
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
			// Set image url for place
			if gptResponse.ImageURL == "" {
				imageURL, err := svr.imgSvc.GetImageURL(newBookmark.PlaceName)
				if err != nil {
					slog.Error("Unable to get image from pexels", "error", err)
					os.Exit(1)
				}
				gptResponse.ImageURL = imageURL.ImageURL
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
		userBookmarkDetail, err := svr.bookmarkPlaceService.CreateUserBookmark(*user_id, newBookmark.Username, newBookmark.PlaceID, newBookmark.PlaceName, newBookmark.Catagory, newBookmark.PlaceText)
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
		return
	}

	// Console response struct to send back as JSON
	response := map[string]interface{}{
		"message": newBookmark.PlaceName + " has already been bookmarked before",
		"UserID":  user_id,
		"PlaceID": newBookmark.PlaceID,
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
