package users

import (
	"arbitrage-betting/internal/managers/users"
	"github.com/gofiber/fiber/v3"
)

func NewRouter(app *fiber.App, usersManager users.Manager) fiber.Router {
	router := app.Group("/users")

	router.Get("/:id", NewGetUserByIdHandler(usersManager).Serve)
	router.Get("/", NewGetUserByNameHandler(usersManager).Serve)
	router.Post("/", NewAddUserHandler(usersManager).Serve)
	router.Put("/", NewUpdateUserHandler(usersManager).Serve)
	router.Delete("/:id", NewDeleteUserHandler(usersManager).Serve)

	return router
}
