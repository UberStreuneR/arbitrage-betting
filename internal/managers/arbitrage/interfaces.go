package arbitrage

import "arbitrage-betting/internal/models"

type Manager interface {
	GetSummary(int) (models.Summary, error)
}
