package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SearchCountry struct {
	Country  string `json:"country"`
	Catagory string `json:"catagory"`
}

// gptSearchCountry processes the incoming POST request for a country search

func (svr *ApiServer) gptSearchCountry(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON data
	var input SearchCountry

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Process the input data
	fmt.Printf("Received country name: %s, catagory:%s\n", input.Country, input.Catagory)

	// Handle the page and get attractions for the provided country

	response, err := svr.locationSvc.GetDetails(input.Country, input.Catagory)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing search: %s", err), http.StatusInternalServerError)
		return
	}

	// TODO: Bookmark query
	// type PlaceDetailsWithBookmark struct {
	// 	Details     PlaceDetails
	// 	HasBookmark bool
	// }

	// responseData = PlaceDetailsWithBookmark{
	// 	Details:     response,
	// 	HasBookmark: getFromOtherService,
	// }

	// Create the JSON response
	jsData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsData)
	if err != nil {
		http.Error(w, "Failed to write HTML response", http.StatusInternalServerError)
	}
}
