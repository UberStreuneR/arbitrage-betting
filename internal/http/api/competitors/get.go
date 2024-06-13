package competitors

import (
	"arbitrage-betting/internal/models"
	"github.com/gofiber/fiber/v3"
)

func GetCompetitors(c fiber.Ctx) error {
	data := []models.Competitor{
		{"1", "dummy competitor 1"},
		{"2", "dummy competitor 2"},
	}

	return c.JSON(data)
}
