package contract

import "github.com/trng-tr/customer-microservice/internal/infrastructure/out/models"

/* CustomerRepository interface extends  interfaces:
XXXSaveAndReadOneORepository,XXXReadAllORepository,
XXXUpdateORepository, XXXDeleteOneORepository, methods:
SaveO, FindOByID,FindAllO,UpdateO,DeleteO
*/
type CustomerRepository interface {
	XXXSaveAndReadOneORepository[models.Customer, int64]
	XXXReadAllORepository[models.Customer]
	XXXUpdateORepository[models.Customer, int64]
	XXXDeleteOneORepository[models.Customer, int64]
}

/* AddressRepository interface extends  interfaces:
XXXSaveAndReadOneORepository,XXXReadAllORepository,
XXXDeleteOneORepository: methods SaveO, FindOByID,
FindAllO& DeleteO
*/
type AddressRepository interface {
	XXXSaveAndReadOneORepository[models.Address, int64]
	XXXReadAllORepository[models.Address]
	XXXDeleteOneORepository[models.Customer, int64]
}
