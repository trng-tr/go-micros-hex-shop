package in

import (
	"context"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

// CustomerService port d'entrée exposé par l'app à l'exterieur
type InCustomerService interface {
	CreateCustomer(ctx context.Context, bsCustomer domain.BusinessCustomer) (domain.BusinessCustomer, error)
	GetCustomerByID(ctx context.Context, id int64) (domain.BusinessCustomer, error)
	GetAllCustomers(ctx context.Context) ([]domain.BusinessCustomer, error)
	PatchCustomer(ctx context.Context, id int64, patchCustomer domain.PatchBusinessCustomer) (domain.BusinessCustomer, error)
	DeleteCustomer(ctx context.Context, id int64) error
}
