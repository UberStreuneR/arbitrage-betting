package models

type Summary []Arbitrage

type Arbitrage struct {
	Match         Match              `json:"match"`
	BestOdds      map[string]BestOdd `json:"best_odds"`
	MarginPercent float32            `json:"margin_percent"`
}

type BestOdd struct {
	Odds      Odds      `json:"odds"`
	Bookmaker Bookmaker `json:"bookmaker"`
}
