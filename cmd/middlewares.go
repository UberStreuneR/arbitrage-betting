package main

import (
	"arbitrage-betting/internal/models"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v3"
)

func setupMiddlewares(app *fiber.App, config models.AuthConfig) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Secret)},
	}))
}
