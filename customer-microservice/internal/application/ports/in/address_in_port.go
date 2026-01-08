package in

import (
	"context"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

// AddressService port d'entrée exposé que l'app expose à l'exterieur
type AddressService interface {
	CreateAddress(ctx context.Context, address domain.BusinessAddress) (domain.BusinessAddress, error)
	GetAddressByID(ctx context.Context, id int64) (domain.BusinessAddress, error)
}
