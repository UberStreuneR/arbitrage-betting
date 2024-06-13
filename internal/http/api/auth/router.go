package auth

import (
	"arbitrage-betting/internal/managers/auth"
	"arbitrage-betting/internal/models"
	"github.com/gofiber/fiber/v3"
)

func NewRouter(app *fiber.App, config models.AuthConfig, authManager auth.Manager) fiber.Router {
	router := app.Group("/auth")

	router.Post("/", NewAuthHandler(config, authManager).Serve)

	return router
}
