package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	EnvName              string `env:"ENV_NAME" envDefault:"local"`
	Version              string `env:"APP_VERSION" envDefault:"local"`
	WebPort              uint16 `env:"WEB_PORT" envDefault:"8080"`
	GptKey               string `env:"OPENAI_API_KEY,required"`
	PsqlConnectionString string `env:"PSQL_CONNECTION_STRING,required"`
}

func startupGetEnv() (*EnvConfig, error) {
	err := godotenv.Load()
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
