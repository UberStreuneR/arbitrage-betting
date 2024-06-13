package bookmakers

import "arbitrage-betting/internal/models"

type Repository interface {
	GetById(string) (models.Bookmaker, error)
	GetByName(string) (models.Bookmaker, error)
}
