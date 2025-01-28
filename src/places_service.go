package main

import (
	"context"
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
