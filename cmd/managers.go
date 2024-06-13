package main

import (
	"arbitrage-betting/internal/managers/arbitrage"
	"arbitrage-betting/internal/managers/auth"
	"arbitrage-betting/internal/managers/matches"
	"arbitrage-betting/internal/managers/odds"
	"arbitrage-betting/internal/managers/users"
)

type managers struct {
	users     users.Manager
	auth      auth.Manager
	matches   matches.Manager
	odds      odds.Manager
	arbitrage arbitrage.Manager
}

func setupManagers(repositories *repositories) managers {
	usersManager := users.NewUserManager(repositories.users)
	matchesManager := matches.NewMatchesManager()
	oddsManager := odds.NewOddsManager()

	return managers{
		users:     usersManager,
		auth:      auth.NewAuthManager(),
		matches:   matchesManager,
		odds:      oddsManager,
		arbitrage: arbitrage.NewArbitrageManager(matchesManager, oddsManager),
	}
}
