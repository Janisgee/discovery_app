package main

import (
	"database/sql"
	"discoveryapp/internal/database"
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

	slog.Info("Successfully connected to the database!")

	// ChatGPT search
	gptClient := openai.NewClient(env.GptKey)
	var locationSvc LocationService = &GptLocationService{gptClient}
	server := NewApiServer(env, locationSvc, userSvc)

	// Start the API server
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
