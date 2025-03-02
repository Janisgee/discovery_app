package main

import (
	"flag"
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
}

func startupGetEnv() (*EnvConfig, error) {
	envPath := flag.String("env", ".env", "path to .env file")
	flag.Parse()

	err := godotenv.Load(*envPath)
	if err != nil {
		return nil, err
	}

	envCfg := EnvConfig{}

	err = env.Parse(&envCfg)
	if err != nil {
		return nil, err
	}

	return &envCfg, nil
}
