package bookmakers

import "github.com/gofiber/fiber/v3"

func NewRouter(app *fiber.App) fiber.Router {
	router := app.Group("/bookmakers")

	// router.Get("/all", GetCompetitors)

	return router
}
