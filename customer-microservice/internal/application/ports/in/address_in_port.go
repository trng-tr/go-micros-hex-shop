package in

import (
	"context"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

// AddressService port d'entrée exposé par l'app à l'exterieur
type InAddressService interface {
	CreateAddress(ctx context.Context, bsAddress domain.BusinessAddress) (domain.BusinessAddress, error)
	GetAddressByID(ctx context.Context, id int64) (domain.BusinessAddress, error)
	GetAllAddresses(ctx context.Context) ([]domain.BusinessAddress, error)
	DeleteAddress(ctx context.Context, id int64) error
}
