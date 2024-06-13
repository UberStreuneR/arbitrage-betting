package bookmakers

import "arbitrage-betting/internal/models"

type Repository interface {
	GetOddsForMatch(models.Match) ([]models.BookmakerOdds, error)
	GetOddsForMatchByBookmaker(models.Match, models.Bookmaker) (models.BookmakerOdds, error)
}
