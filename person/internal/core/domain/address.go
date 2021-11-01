package domain

type Address struct {
	ID         string  `json:"id"`
	Country    Country `json:"country"`
	State      string  `json:"state"`
	City       string  `json:"city"`
	PostCode   string  `json:"postCode"`
	Address    string  `json:"address"`
	Number     string  `json:"number"`
	Complement string  `json:"complement"`

	Person Person `json:"person"`
}
