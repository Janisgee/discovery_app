package main

import (
	"database/sql"
	"discoveryapp/internal/database"
	"fmt"
	"log"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/sashabaranov/go-openai"
)

func main() {

	// Setup logger
	jsonLogs := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(jsonLogs)

	// Load environment file
	env, err := startupGetEnv()
	if err != nil {
		slog.Error("error loading environment config", "error", err)
		os.Exit(1)
	}

	// Connect to PSQL database
	// Open a connection to the database
	db, err := sql.Open("postgres", env.PsqlConnectionString)
	if err != nil {
		slog.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	// Use generated database package to create a new *database.Queries instance
	dbQueries := database.New(db)

	// Create user service to run the queries
	var userSvc UserService = &PostgresUserService{dbQueries}

	slog.Info("Successfully connected to the database!", "Connection String", dbQueries)

	// Connect to pexels image client
	imageSvc := NewPexelsService(env.PexelsKey)
	result, err := imageSvc.GetImageURL("Tokyo")
	if err != nil {
		slog.Error("Unable to get image from pexels", "error", err)
		os.Exit(1)
	}
	fmt.Println(result)

	// ChatGPT search
	gptClient := openai.NewClient(env.GptKey)
	var locationSvc LocationService = &GptLocationService{gptClient}

	// Google place service
	placesSvc, err := NewGooglePlacesService(env.GMapsKey)
	if err != nil {
		slog.Error("Unable to connect google maps service", "error", err)
		os.Exit(1)
	}

	// Create bookmark place service to run the queries
	var bookmarkPlaceSvc BookmarkPlaceService = &PostgresBookmarkService{dbQueries}

	server := NewApiServer(env, locationSvc, userSvc, placesSvc, bookmarkPlaceSvc, imageSvc)

	// Start the API server
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
