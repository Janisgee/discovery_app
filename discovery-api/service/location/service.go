package location

import (
	"context"
	"database/sql"
	"discoveryweb/internal/database"
	"discoveryweb/service/places"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

type gptLocationService struct {
	gptClient     *openai.Client
	placesService places.PlacesService
	dbQueries     *database.Queries
}

func NewGptService(client *openai.Client, placesService places.PlacesService, dbQueries *database.Queries) LocationService {
	return &gptLocationService{
		gptClient:     client,
		placesService: placesService,
		dbQueries:     dbQueries,
	}
}

func (svc *gptLocationService) GetDetails(location string, category string) ([]LocationDetails, error) {
	prompt := fmt.Sprintf("Get 3 %s in %s, using a field 'places' containing  'name' (the attraction name) and 'description' (a 10-word description).", category, location)

	completion, err := svc.gptClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant and response in json object format without ```json"},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		slog.Warn("Failure on chat gpt completion")
		return nil, fmt.Errorf("ChatCompletion error: %v\n ", err)
	}

	content := completion.Choices[0].Message.Content

	var response struct {
		Places []LocationDetails `json:"places"`
	}

	err = json.Unmarshal([]byte(content), &response)
	if err != nil {
		return nil, fmt.Errorf("ChatGpt json error: %v\n ", err)
	}

	// Inject place_id field manually
	for i := range response.Places {

		if response.Places[i].PlaceID == "" { // Check if place_id is missing
			// Encoded space to %20
			encodeLocation := strings.ReplaceAll(response.Places[i].Name, " ", "%20")

			// Search place place id from Google map
			place_id, err := svc.placesService.GetPlaceID(encodeLocation)
			if err != nil {
				slog.Warn("Error getting place id for the search place from Google map in GetPlaceDetail")
				return nil, fmt.Errorf("error getting place id for the search place from Google map in GetPlaceDetail: %v\n ", err)
			}
			response.Places[i].PlaceID = place_id

		}
	}

	return response.Places, nil
}

func (svc *gptLocationService) GetPlaceDetails(location string) (*PlaceDetails, error) {
	prompt := fmt.Sprintf("Get details of %s, using a field 'place_details' containing 'city'(which city %s belong to, if no detail of city, which region instead),'country'(which country %s belong to),'description'(around 20 words),'location' (address), 'opening_hours' (everyday operation hour), 'history' (around 50 words), 'key_features' (around 100 words) and 'conclusion'(around 40 words conclusion for the place).", location, location, location)

	completion, err := svc.gptClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant and response in json object format without ```json"},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		slog.Warn("Failure on chat gpt completion")
		return nil, fmt.Errorf("ChatCompletion error: %v\n ", err)
	}
	content := completion.Choices[0].Message.Content

	slog.Debug("Raw GPT response", "content", content)

	var response struct {
		Place_details PlaceDetails `json:"place_details"`
	}

	err = json.Unmarshal([]byte(content), &response)
	if err != nil {
		return nil, fmt.Errorf("ChatGpt json error: %v\n ", err)
	}

	return &response.Place_details, nil
}

func (svc *gptLocationService) CheckCountryCityInData(country string, city string) (string, error) {

	// Create an empty context
	ctx := context.Background()

	// Check if queries country is in database
	_, err := svc.dbQueries.GetCountry(ctx, country)
	if err != nil {
		slog.Warn("Fail to get country from database", "error", err)
		return "No country image", errors.New("fail to get country from database")
	}

	// Check if queries city is in database
	params := database.GetCityParams{
		City:    city,
		Country: country,
	}
	_, err = svc.dbQueries.GetCity(ctx, params)
	if err != nil {
		slog.Warn("Fail to get city from database", "error", err)
		return "No city image", errors.New("fail to get city from database")
	}

	return "With country and city image", nil
}

func (svc *gptLocationService) GetCountryImageData(country string) (*CountryImage, error) {

	// Create an empty context
	ctx := context.Background()

	// Check if queries country is in database
	countryImageData, err := svc.dbQueries.GetCountry(ctx, country)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No rows found, return specific error for "not found"
			slog.Warn("No rows found from country image database", "error", err)
			return nil, errors.New("no rows found from country image database")
		}
		slog.Warn("Fail to get country from database", "error", err)
		return nil, errors.New("fail to get country from database")
	}

	// Return user bookmark place
	data := &CountryImage{
		CountryID:    countryImageData.ID,
		Country:      countryImageData.Country,
		CountryImage: countryImageData.CountryImage,
	}

	return data, nil
}

func (svc *gptLocationService) CreateCityImageData(countryID uuid.UUID, country string, city string, cityImage string) (*CityImage, error) {
	// Create an empty context
	ctx := context.Background()
	// Check if queries city is in database
	params := database.CreateCityImageParams{
		CountryID: countryID,
		Country:   country,
		City:      city,
		CityImage: cityImage,
	}
	newCityImageData, err := svc.dbQueries.CreateCityImage(ctx, params)
	if err != nil {
		slog.Warn("Fail to store image into cityImage", "error", err)
		return nil, errors.New("fail to store image into cityImage")
	}

	// Return user bookmark place
	data := &CityImage{
		City:      newCityImageData.City,
		CityImage: newCityImageData.CityImage,
	}

	return data, nil

}

func (svc *gptLocationService) CreateCountryImageData(country string, countryImage string) (*CountryImage, error) {
	// Create an empty context
	ctx := context.Background()
	// Check if queries country is in database
	params := database.CreateCountryImageParams{
		Country:      country,
		CountryImage: countryImage,
	}
	newCountryImageData, err := svc.dbQueries.CreateCountryImage(ctx, params)
	if err != nil {
		slog.Warn("Fail to store image into countryImage database", "error", err)
		return nil, errors.New("fail to store image into countryImage database")
	}

	// Return user bookmark place
	data := &CountryImage{
		CountryID:    newCountryImageData.ID,
		Country:      newCountryImageData.Country,
		CountryImage: newCountryImageData.CountryImage,
	}

	return data, nil

}
