package bookmakers

import "arbitrage-betting/internal/models"

type Repository interface {
	GetById(string) (models.Competitor, error)
	GetByName(string) (models.Competitor, error)
}
