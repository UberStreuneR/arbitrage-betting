package arbitrage

import (
	"arbitrage-betting/internal/managers/auth"
	"arbitrage-betting/internal/models"
	"github.com/gofiber/fiber/v3"
)

func NewArbitrageHandler(config models.AuthConfig, authManager auth.Manager) *arbitrageHandler {
	return &arbitrageHandler{config: config, authManager: authManager}
}

type arbitrageHandler struct {
	config      models.AuthConfig
	authManager auth.Manager
}

func (h *arbitrageHandler) Serve(c fiber.Ctx) error {
	chelsea := models.Competitor{
		ID:   "chelsea_id",
		Name: "Chelsea",
	}

	arsenal := models.Competitor{
		ID:   "arsenal_id",
		Name: "Arsenal",
	}

	goodbets := models.Bookmaker{
		ID:   "1",
		Name: "GoodBets",
	}

	xbet := models.Bookmaker{
		ID:   "2",
		Name: "1xBet",
	}

	summary := models.Summary{
		{
			Match: models.Match{
				ID:          "1",
				Competitor1: chelsea,
				Competitor2: arsenal,
				Completed:   false,
			},
			BestOdds: map[string]models.BestOdd{
				"chelsea_id": {
					Odds: models.Odds{
						Competitor:  chelsea,
						Coefficient: 1.56,
					},
					Bookmaker: goodbets,
				},
				"arsenal_id": {
					Odds: models.Odds{
						Competitor:  arsenal,
						Coefficient: 2.87,
					},
					Bookmaker: xbet,
				},
			},
			MarginPercent: 1.07,
		},
	}

	return c.Status(fiber.StatusOK).JSON(summary)
}
