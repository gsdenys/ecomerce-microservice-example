package domain

type City struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Region Region `json:"region"`
}
