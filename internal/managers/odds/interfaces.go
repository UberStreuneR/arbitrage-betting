package odds

import "arbitrage-betting/internal/models"

type Manager interface {
	GetOddsForMatch(models.Match) ([]models.BookmakerOdds, error)
	GetOddsForMatchByBookmaker(models.Match, models.Bookmaker) (models.BookmakerOdds, error)
}

func NewOddsManager() *oddsManager {
	return &oddsManager{}
}

type oddsManager struct{}

func (oddsManager) GetOddsForMatch(match models.Match) ([]models.BookmakerOdds, error) {
	// TODO implement me
	panic("implement me")
}

func (oddsManager) GetOddsForMatchByBookmaker(match models.Match, bookmaker models.Bookmaker) (models.BookmakerOdds, error) {
	// TODO implement me
	panic("implement me")
}
