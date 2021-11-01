package countrysrv

import (
	"github.com/codemodus/uidgen"
	"github.com/gsdenys/ecomerce-microservice-example/address/internal/core/port"
)

type service struct {
	countryRepository port.CountryRepository
	uidGen            uidgen.UIDGen
}
