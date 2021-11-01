package port

import "github.com/gsdenys/ecomerce-microservice-example/address/internal/core/domain"

type CountryRepository interface {
	Get(id string) (domain.Country, error)
	Save(domain.Country) error
}

type CountryService interface {
	Get(id string) (domain.Country, error)
	Create(name string, acronym string) (domain.Country, error)
}
