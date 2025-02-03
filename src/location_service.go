package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/sashabaranov/go-openai"
)

type LocationService interface {
	GetDetails(location string, category string) ([]LocationDetails, error)
	GetPlaceDetails(location string) (*PlaceDetails, error)
}

type LocationDetails struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	PlaceID     string `json:"place_id"`
	Description string `json:"description"`
}

type PlaceDetails struct {
	City          string `json:"city"`
	Country       string `json:"country"`
	Description   string `json:"description"`
	Location      string `json:"location"`
	Opening_hours string `json:"opening_hours"`
	History       string `json:"history"`
	Key_features  string `json:"key_features"`
	Conclusion    string `json:"conclusion"`
}

type GptLocationService struct {
	gptClient *openai.Client
}

func (svc *GptLocationService) GetDetails(location string, category string) ([]LocationDetails, error) {
	prompt := fmt.Sprintf("Get 3 %s in %s, using a field 'places' containing 'image' (a URL to an image), 'name' (the attraction name), 'place_id'(google place id of %s)  and 'description' (a 10-word description).", category, location, location)

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

	return response.Places, nil
}

func (svc *GptLocationService) GetPlaceDetails(location string) (*PlaceDetails, error) {
	prompt := fmt.Sprintf("Get details of %s, using a field 'place_details' containing 'city'(which city %s belong to),'country'(which country %s belong to),'description'(around 20 words),'location' (address), 'opening_hours' (everyday operation hour), 'history' (around 50 words), 'key_features' (around 100 words) and 'conclusion'(around 40 words conclusion for the place).", location, location, location)

	// fmt.Printf("GetPlaceDetails location: %s \n")
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
