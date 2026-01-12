package out

import (
	"context"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

// OutAddressService port de sortie utilisé par l'app pour envoyer à l'exterieur
type OutAddressService interface {
	CreateAddress(ctx context.Context, bsAddress domain.BusinessAddress) (domain.BusinessAddress, error)
	GetAddressByID(ctx context.Context, id int64) (domain.BusinessAddress, error)
	GetAllAddresses(ctx context.Context) ([]domain.BusinessAddress, error)
	DeleteAddress(ctx context.Context, id int64) error
}
