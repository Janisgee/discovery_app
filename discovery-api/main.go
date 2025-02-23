package main

import (
	"database/sql"
	"discoveryweb/api"
	"discoveryweb/internal/database"
	"discoveryweb/service/bookmark"
	"discoveryweb/service/email"
	"discoveryweb/service/image"
	"discoveryweb/service/location"
	"discoveryweb/service/places"
	"discoveryweb/service/user"
	"fmt"
	gomail "gopkg.in/mail.v2"
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
	slog.Info("Successfully connected to the database!")

	defer db.Close()

	err = migrateDb(db)
	if err != nil {
		slog.Error("Unable to complete goose migrations", "err", err)
		os.Exit(1)
	}

	// Use generated database package to create a new *database.Queries instance
	dbQueries := database.New(db)

	// Create user service to run the queries
	var userSvc = user.NewUserService(dbQueries)

	// Connect to pexels image client
	imageSvc := image.NewPexelsService(env.PexelsKey)
	result, err := imageSvc.GetImageURL("Tokyo")
	if err != nil {
		slog.Error("Unable to get image from pexels", "error", err)
		os.Exit(1)
	}
	fmt.Println(result)

	// Google place service
	placesSvc, err := places.NewGooglePlacesService(env.GMapsKey)
	if err != nil {
		slog.Error("Unable to connect google maps service", "error", err)
		os.Exit(1)
	}

	// ChatGPT search
	gptClient := openai.NewClient(env.GptKey)
	var locationSvc = location.NewGptService(gptClient, placesSvc)

	// Create bookmark place service to run the queries
	var bookmarkPlaceSvc = bookmark.NewBookmarkPlaceService(dbQueries)

	var mailDialer = gomail.NewDialer("smtp.gmail.com", 587, "yanisching@gmail.com", "yjpyfqdkwkczydef")
	var emailSvc = email.NewEmailService("yanisching@gmail.com", mailDialer)

	server := api.NewApiServer(env.WebPort, locationSvc, userSvc, placesSvc, bookmarkPlaceSvc, imageSvc, emailSvc)

	// Start the API server
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
