package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strings"
)

type SearchPlaceDetails struct {
	Place    string `json:"place"`
	Catagory string `json:"catagory"`
}

// gptSearchPlaceDetails processes the incoming POST request for a place search

func (svr *ApiServer) gptSearchPlaceDetails(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("Received place name (line34): %s \n", input.Place)

	// Search place place id from Google map
	place_id, err := getPlaceID(input.Place)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting place id for the %s from Google map:%s", input.Place, err), http.StatusInternalServerError)
		slog.Error("error getting place id for the search place from Google map", "error", err)
		return
	}
	// Encoded space to %20
	encodeLocation := strings.ReplaceAll(input.Place, "%20", " ")

	// Check if place has been stored in database
	placeInDB, err := svr.bookmarkPlaceService.GetPlaceDatabaseDetails(place_id)
	if err != nil {
		// [NO] No place found from place id in Database

		// Handle the page and get place details for the provided place from ChatGPT
		response, err := svr.locationSvc.GetPlaceDetails(encodeLocation)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting place detail: %s", err), http.StatusInternalServerError)
			slog.Error("Error getting place detail", "error", err)
			return
		}

		// Store place into place Data
		err = svr.bookmarkPlaceService.CreatePlaceData(place_id, encodeLocation, response.Country, response.City, input.Catagory, response)
		if err != nil {
			http.Error(w, "Failed to create place into place database.", http.StatusBadRequest)
			slog.Error("failed to create place into place database", "error", err)
			return
		}

		fmt.Println("Go through CHATGPT to see place detail")

		////////////////////////Get bookmark check//////////////////
		// responseWithBookmarkInfo struct
		type responseWithBookmarkInfo struct {
			PlaceInfo   *PlaceDetails
			HasBookmark bool
		}
		// Get user detail
		user_id := GetCurrentUserId(r)

		// Check if place has been bookmarked by user (Return true or false)
		result, _ := svr.bookmarkPlaceService.CheckPlaceHasBookmarkedByUser(place_id, *user_id)

		// Inject response struct with bookmark check
		finalResponse := responseWithBookmarkInfo{
			PlaceInfo:   response,
			HasBookmark: result,
		}
		////////////////////////Finish bookmark check//////////////////
		// Create the JSON response
		jsData, err := json.Marshal(finalResponse)
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
		fmt.Println("Store data into database through CHATGPT")
		return
	}

	// [YES]
	fmt.Println("Go through database to see place detail")

	////////////////////////Get bookmark check//////////////////
	// responseWithBookmarkInfo struct
	type responseWithBookmarkInfo struct {
		PlaceInfo   PlaceDetails
		HasBookmark bool
	}
	// Get user detail
	user_id := GetCurrentUserId(r)

	// Check if place has been bookmarked by user (Return true or false)
	result, _ := svr.bookmarkPlaceService.CheckPlaceHasBookmarkedByUser(place_id, *user_id)

	// Inject response struct with bookmark check
	finalResponse := responseWithBookmarkInfo{
		PlaceInfo:   placeInDB.PlaceDetail,
		HasBookmark: result,
	}
	////////////////////////Finish bookmark check//////////////////
	// Create the JSON response
	jsData, err := json.Marshal(finalResponse)
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
