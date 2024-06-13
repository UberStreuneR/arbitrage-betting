package models

type Odds struct {
	Competitor  Competitor `json:"competitor"`
	Coefficient float32    `json:"coefficient"`
}

type BookmakerOdds struct {
	Bookmaker Bookmaker `json:"bookmaker"`
	Match     Match     `json:"match"`
	Odds      []Odds    `json:"odds"`
}
