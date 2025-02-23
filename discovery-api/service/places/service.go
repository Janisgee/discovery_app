package places

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/places/v1"
	"log/slog"
	"net/http"
)

type googlePlacesService struct {
	ggPlacesClient *places.Service
	key            string
}

// Create a google client
func NewGooglePlacesService(key string) (PlacesService, error) {
	service, err := places.NewService(context.Background(), option.WithAPIKey(key))
	if err != nil {
		return nil, err
	}
	return &googlePlacesService{
		ggPlacesClient: service,
		key:            key,
	}, nil
}

func (svc *googlePlacesService) AutocompleteCities(search string, lang string) ([]CityResult, error) {
	result, err := svc.ggPlacesClient.Places.Autocomplete(&places.GoogleMapsPlacesV1AutocompletePlacesRequest{
		IncludedPrimaryTypes: []string{"administrative_area_level_1", "administrative_area_level_2", "country"},
		LanguageCode:         lang,
		Input:                search,
	}).Do()

	if err != nil {
		return nil, err
	}

	testing, err := svc.ggPlacesClient.Places.Get("places/ChIJpU8j7H-1HGARxU4d9u5v9qA").Fields("id", "displayName").Do()
	if err != nil {
		slog.Error("failed to get place by id", "error", err)
	}
	slog.Info("Places result", "place", testing)

	cities := make([]CityResult, len(result.Suggestions))

	for ii, suggestion := range result.Suggestions {
		cities[ii] = CityResult{
			Id:   suggestion.PlacePrediction.PlaceId,
			Name: suggestion.PlacePrediction.StructuredFormat.MainText.Text,
		}
		if suggestion.PlacePrediction.StructuredFormat.SecondaryText == nil {
			cities[ii].Region = suggestion.PlacePrediction.StructuredFormat.MainText.Text
		} else {
			cities[ii].Region = suggestion.PlacePrediction.StructuredFormat.SecondaryText.Text
		}
	}

	return cities, nil
}

func (svc *googlePlacesService) GetPlaceID(location string) (string, error) {
	// Get key
	// Construct the URL for the Google Places API FindPlaceFromText request
	baseURL := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json"
	requestURL := fmt.Sprintf("%s?input=%s&inputtype=textquery&fields=place_id&key=%s", baseURL, location, svc.key)
	fmt.Printf("RequestURL to google api: %s", requestURL)
	// Send the HTTP GET request to Google Places API
	/* #nosec */
	resp, err := http.Get(requestURL)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}

	defer resp.Body.Close()

	// Decode the JSON response
	var response placeFindResponse
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
