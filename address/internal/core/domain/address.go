package domain

type Address struct {
	ID         string  `json:"id"`
	Country    Country `json:"country"`
	Region     Region  `json:"region"`
	City       City    `json:"city"`
	PostCode   string  `json:"postcode"`
	Address    string  `json:"address"`
	Number     string  `json:"number"`
	Complement string  `json:"complement"`
}
