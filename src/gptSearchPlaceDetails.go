package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
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

	// Search place place id from ChatGPT
	place_id, err := getPlaceID(input.Place)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting place id for the %s from ChatGPT:%s", input.Place, err), http.StatusInternalServerError)
		slog.Error("error getting place id for the search place from ChatGPT", "error", err)
		return
	}

	fmt.Println("GetPlaceDatabaseDetails Helllllllo place_id:", place_id)
	// Check if place has been stored in database
	placeInDB, err := svr.bookmarkPlaceService.GetPlaceDatabaseDetails(place_id)
	if err != nil {
		// [NO] No place found from place id in Database

		// Handle the page and get place details for the provided place from ChatGPT
		response, err := svr.locationSvc.GetPlaceDetails(input.Place)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting place detail: %s", err), http.StatusInternalServerError)
			slog.Error("Error getting place detail", "error", err)
			return
		}

		fmt.Println("CreatePlaceData: placeID", place_id)
		fmt.Println("CreatePlaceData: place name", input.Place)
		fmt.Println("CreatePlaceData: place country", response.Country)
		fmt.Println("CreatePlaceData: place city", response.City)
		fmt.Println("CreatePlaceData: catagory", input.Catagory)
		fmt.Println("CreatePlaceData: place details", response)

		// Store place into place Data
		err = svr.bookmarkPlaceService.CreatePlaceData(place_id, input.Place, response.Country, response.City, input.Catagory, response)
		if err != nil {
			http.Error(w, "Failed to create place into place database.", http.StatusBadRequest)
			slog.Error("failed to create place into place database", "error", err)
			return
		}

		fmt.Println("XXXXXXXXXXXXXXXXXXXX Go through CHATGPT to see place detail")
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
		fmt.Println(" SSSSSStore data into database through CHATGPT ")
		return
	}

	// [YES]
	fmt.Println("XXXXXXXXXXXXXXXXXXXX Go through database to see place detail")
	// Create the JSON response
	jsData, err := json.Marshal(placeInDB.PlaceDetail)
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
