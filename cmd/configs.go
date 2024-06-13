package main

import (
	"arbitrage-betting/internal/models"
	"os"
)

const (
	authSecretEnvVariable = "SECRET_KEY"
)

type configs struct {
	auth models.AuthConfig
}

func setupConfigs() configs {
	authConfig := models.AuthConfig{
		Secret: os.Getenv(authSecretEnvVariable),
	}

	return configs{
		auth: authConfig,
	}
}
