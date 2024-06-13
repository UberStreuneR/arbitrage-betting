package matches

import "arbitrage-betting/internal/models"

type Manager interface {
	GetById(string) (models.Match, error)
	GetAll(string) ([]models.Match, error)
	GetMostRecent(int) ([]models.Match, error)
	AddMatch(models.Match) error
}

func NewMatchesManager() *matchesManager {
	return &matchesManager{}
}

type matchesManager struct{}

func (matchesManager) GetById(s string) (models.Match, error) {
	// TODO implement me
	panic("implement me")
}

func (matchesManager) GetAll(s string) ([]models.Match, error) {
	// TODO implement me
	panic("implement me")
}

func (matchesManager) GetMostRecent(i int) ([]models.Match, error) {
	// TODO implement me
	panic("implement me")
}

func (matchesManager) AddMatch(match models.Match) error {
	// TODO implement me
	panic("implement me")
}
