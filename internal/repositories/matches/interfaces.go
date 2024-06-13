package bookmakers

import "arbitrage-betting/internal/models"

type Repository interface {
	GetById(string) (models.Match, error)
	GetAll(string) ([]models.Match, error)
	GetMostRecent(int) ([]models.Match, error)
}
