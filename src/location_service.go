package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log/slog"
)

type LocationService interface {
	GetDetails(location string, category string) ([]LocationDetails, error)
}

type LocationDetails struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GptLocationService struct {
	gptClient *openai.Client
}

func (svc *GptLocationService) GetDetails(location string, category string) ([]LocationDetails, error) {
	prompt := fmt.Sprintf("Get 3 %s in %s, using a field 'places' containing 'image' (a URL to an image), 'name' (the attraction name), and 'description' (a 10-word description).", category, location)

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
		return nil, fmt.Errorf("ChatGpt json error: %v\n", err)
	}

	return response.Places, nil
}
