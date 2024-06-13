package main

import (
	"arbitrage-betting/internal/http/api/arbitrage"
	"arbitrage-betting/internal/http/api/auth"
	"arbitrage-betting/internal/http/api/bookmakers"
	"arbitrage-betting/internal/http/api/competitors"
	"arbitrage-betting/internal/http/api/users"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.Get("/healthcheck", func(c fiber.Ctx) error {
		return c.Send([]byte("Hello world"))
	})

	configs := setupConfigs()
	repositories := setupRepositories()
	managers := setupManagers(repositories)

	// setupMiddlewares(app, configs.auth)
	addRouters(app, managers, configs)

	app.Listen(":8000")
}

func addRouters(
	app *fiber.App,
	managers managers,
	configs configs,
) {
	auth.NewRouter(app, configs.auth, managers.auth)
	competitors.NewRouter(app)
	users.NewRouter(app, managers.users)
	bookmakers.NewRouter(app)
	arbitrage.NewRouter(app, configs.auth, managers.auth)
}
