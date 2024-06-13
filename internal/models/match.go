package models

type Match struct {
	ID          string     `json:"id"`
	Competitor1 Competitor `json:"competitor_1"`
	Competitor2 Competitor `json:"competitor_2"`
	Completed   bool       `json:"completed"`
}
