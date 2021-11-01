package domain

type Region struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Acronym string  `json:"acronym"`
	Country Country `json:"country"`
}
