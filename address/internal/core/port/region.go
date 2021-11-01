package port

import "github.com/gsdenys/ecomerce-microservice-example/address/internal/core/domain"

type RegionRepository interface {
	Get(id string) (domain.Region, error)
	List(countryID string) ([]domain.Region, error)
	Save(domain.Region) error
}

type RegionService interface {
	Get(id string) (domain.Region, error)
	Create(name string, acronym string, country domain.Country) (domain.Region, error)
}
