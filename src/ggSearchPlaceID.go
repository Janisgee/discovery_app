package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

// Define the response structure that matches the API's JSON response format
type PlaceFindResponse struct {
	Candidates []struct {
		PlaceID string `json:"place_id"`
	} `json:"candidates"`
	Status string `json:"status"`
}

func getPlaceID(location string) (string, error) {
	// Get key
	env, err := startupGetEnv()
	if err != nil {
		slog.Error("error loading environment config", "error", err)
		os.Exit(1)
	}
	// Construct the URL for the Google Places API FindPlaceFromText request
	baseURL := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json"
	escapedQuery := url.QueryEscape(location)

	requestURL := fmt.Sprintf("%s?input=%s&inputtype=textquery&fields=place_id&key=%s", baseURL, escapedQuery, env.GMapsKey)

	// Send the HTTP GET request to Google Places API
	resp, err := http.Get(requestURL)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}

	defer resp.Body.Close()

	// Decode the JSON response
	var response PlaceFindResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	// If no candidates are found, return an error
	if len(response.Candidates) == 0 {
		return "", fmt.Errorf("no place found for location: %s", location)
	}

	// Return the Place ID from the first candidate result
	return response.Candidates[0].PlaceID, nil

}
