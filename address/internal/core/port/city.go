package port

import "github.com/gsdenys/ecomerce-microservice-example/address/internal/core/domain"

type CityRepository interface {
	Get(id string) (domain.City, error)
	List(regionID string) ([]domain.City, error)
	Save(domain.City) error
}

type CityService interface {
	Get(id string) (domain.City, error)
	Create(name string, acronym string, region domain.Region) (domain.City, error)
}
