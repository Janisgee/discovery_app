package main

import (
	"flag"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	EnvName              string `env:"ENV_NAME" envDefault:"local"`
	Version              string `env:"APP_VERSION" envDefault:"local"`
	WebPort              uint16 `env:"WEB_PORT" envDefault:"8080"`
	GptKey               string `env:"OPENAI_API_KEY,required"`
	GMapsKey             string `env:"GMAPS_API_KEY,required"`
	PsqlConnectionString string `env:"PSQL_CONNECTION_STRING,required"`
	PexelsKey            string `env:"PEXEL_API_KEY,required"`
	ClientBaseUrl        string `env:"CLIENT_BASE_URL,required"`
	GgmailerKey          string `env:"GGMAILER_KEY"`
	GgmailerEmail        string `env:"GGMAILER_EMAIL"`
}

func startupGetEnv() (*EnvConfig, error) {
	envPath := flag.String("env", ".env", "path to .env file")
	flag.Parse()

	//Skip .env Loading on Render: Modify your startupGetEnv() function to only load the .env file if you are running the application locally. If ENV_NAME is empty, it's assumed to be local, and the .env file is loaded.
	if os.Getenv("IS_RENDER") == "" {
		err := godotenv.Load(*envPath)
		if err != nil {
			return nil, err
		}
	}

	envCfg := EnvConfig{}
	err := env.Parse(&envCfg)
	if err != nil {
		return nil, err
	}

	return &envCfg, nil
}
