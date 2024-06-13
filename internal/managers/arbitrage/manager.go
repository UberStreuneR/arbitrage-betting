package arbitrage

import (
	"arbitrage-betting/internal/managers/matches"
	"arbitrage-betting/internal/managers/odds"
	"arbitrage-betting/internal/models"
)

func NewArbitrageManager(matchesManager matches.Manager, oddsManager odds.Manager) Manager {
	return &arbitrageManager{matchesManager: matchesManager, oddsManager: oddsManager}
}

type arbitrageManager struct {
	matchesManager matches.Manager
	oddsManager    odds.Manager
}

func (m *arbitrageManager) GetSummary(numberOfMatches int) (models.Summary, error) {
	var summary models.Summary

	matches, err := m.matchesManager.GetMostRecent(numberOfMatches)
	if err != nil {
		return summary, err
	}

	for _, match := range matches {
		arbitrage, err := m.getArbitrage(match)
		if err != nil {
			return summary, err
		}

		profitMargin := m.getProfitMargin(arbitrage)
		if profitMargin > 0 {
			summary = append(summary, arbitrage)
		}
	}

	return summary, nil
}

func (m *arbitrageManager) getArbitrage(match models.Match) (models.Arbitrage, error) {
	arbitrage := models.Arbitrage{Match: match}

	if match.Completed {
		return models.Arbitrage{}, nil
	}

	bookmakerOdds, err := m.oddsManager.GetOddsForMatch(match)
	if err != nil {
		return models.Arbitrage{}, err
	}

	for _, bookmakerOdd := range bookmakerOdds {
		m.processOdds(bookmakerOdd.Odds, bookmakerOdd.Bookmaker, arbitrage.BestOdds)
	}

	return arbitrage, nil
}

func (m *arbitrageManager) processOdds(
	odds []models.Odds,
	bookmaker models.Bookmaker,
	bestOdds map[string]models.BestOdd,
) {
	for _, odd := range odds {
		id := odd.Competitor.ID
		currentBestOdds, ok := bestOdds[id]
		if !ok {
			bestOdds[id] = models.BestOdd{Odds: odd, Bookmaker: bookmaker}
			continue
		}

		if odd.Coefficient > currentBestOdds.Odds.Coefficient {
			bestOdds[id] = models.BestOdd{Odds: odd, Bookmaker: bookmaker}
		}
	}
}

func (m *arbitrageManager) getProfitMargin(arbitrage models.Arbitrage) float32 {
	coefficient1, coefficient2 := m.getBestOddsFromArbitrage(arbitrage)
	margin := (1/coefficient1 + 1/coefficient2) - 1
	return -margin
}

func (m *arbitrageManager) getBestOddsFromArbitrage(arbitrage models.Arbitrage) (float32, float32) {
	odd1, ok := arbitrage.BestOdds[arbitrage.Match.Competitor1.ID]
	if !ok {
		return 0, 0
	}

	odd2, ok := arbitrage.BestOdds[arbitrage.Match.Competitor2.ID]
	if !ok {
		return 0, 0
	}

	return odd1.Odds.Coefficient, odd2.Odds.Coefficient
}
