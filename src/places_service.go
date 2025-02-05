package main

import (
	"context"
	"log/slog"

	"google.golang.org/api/option"
	"google.golang.org/api/places/v1"
)

type CityResult struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

type PlacesService interface {
	AutocompleteCities(search string, lang string) ([]CityResult, error)
}

type GooglePlacesService struct {
	ggPlacesClient *places.Service
}

// Create a google client
func NewGooglePlacesService(key string) (*GooglePlacesService, error) {
	service, err := places.NewService(context.Background(), option.WithAPIKey(key))
	if err != nil {
		return nil, err
	}
	return &GooglePlacesService{
		ggPlacesClient: service,
	}, nil
}

func (svc *GooglePlacesService) AutocompleteCities(search string, lang string) ([]CityResult, error) {
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
