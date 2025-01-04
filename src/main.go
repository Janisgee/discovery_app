package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"log/slog"
	"os"
)

func main() {
	// Setup logger
	jsonLogs := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(jsonLogs)

	// Load environment file
	env, err := startupGetEnv()
	if err != nil {
		fmt.Printf("error loading environment config :%s \n ", err)
		os.Exit(1)
	}

	gptClient := openai.NewClient(env.GptKey)
	var locationSvc LocationService = &GptLocationService{gptClient}
	server := NewApiServer(env, locationSvc)

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
