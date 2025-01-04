package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// handleSearchCountry processes the incoming POST request for a country search

func handleSearchCountry(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("Received country name: %s\n", input.Country)

	// Handle the page and get attractions for the provided country
	response, err := handleSearch(input.Country)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing search: %s", err), http.StatusInternalServerError)
		return
	}
	handlePage(w, response)
}
