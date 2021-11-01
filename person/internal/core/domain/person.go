package domain

import "time"

type Person struct {
	ID         string    `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	SocialName string    `json:"socialName"`
	BirthDate  time.Time `json:"birthDate"`
}
