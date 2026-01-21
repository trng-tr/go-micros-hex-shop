package mappers

import (
	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
)

func ToDomainCustomer(dtoResp dtos.CustomerResponse) domain.Customer {
	return domain.Customer{
		ID:          dtoResp.ID,
		Firstname:   dtoResp.Firstname,
		Lastname:    dtoResp.Lastname,
		Genda:       domain.Genda(dtoResp.Genda),
		Email:       dtoResp.Email,
		PhoneNumber: dtoResp.PhoneNumber,
		Status:      domain.CustomerStatus(dtoResp.Status),
	}
}

func ToCustomerResponse(customer domain.Customer) dtos.LightCustomerResponse {
	return dtos.LightCustomerResponse{
		ID:          customer.ID,
		Firstname:   customer.Firstname,
		Lastname:    customer.Lastname,
		Genda:       string(customer.Genda),
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Status:      string(customer.Status),
	}
}
