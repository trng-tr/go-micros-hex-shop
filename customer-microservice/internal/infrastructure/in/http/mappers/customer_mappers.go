package mappers

import (
	"github.com/trng-tr/customer-microservice/internal/domain"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/dtos"
)

func ToBusinessCustomer(request dtos.CustomerRequest) domain.BusinessCustomer {
	return domain.BusinessCustomer{
		Firstname:   request.Firstname,
		Lastname:    request.Lastname,
		Genda:       domain.Genda(request.Genda),
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		AddressID:   request.AddressID,
	}
}

func ToCustomerResponse(bs domain.BusinessCustomer, bsAddress domain.BusinessAddress) dtos.CustomerResponse {
	return dtos.CustomerResponse{
		ID:              bs.ID,
		Firstname:       bs.Firstname,
		Lastname:        bs.Lastname,
		Genda:           string(bs.Genda),
		Email:           bs.Email,
		PhoneNumber:     bs.PhoneNumber,
		Status:          string(bs.Status),
		AddressResponse: ToAddressResponse(bsAddress),
		CreatedAt:       bs.CreatedAt,
		UpdatedAt:       bs.UpdatedAt,
	}
}

func ToDomainPatch(patchRequest dtos.CustomerPatchRequest) domain.PatchBusinessCustomer {
	return domain.PatchBusinessCustomer{
		Firstname:   patchRequest.Firstname,
		Lastname:    patchRequest.Lastname,
		Email:       patchRequest.Email,
		PhoneNumber: patchRequest.PhoneNumber,
		AddressID:   patchRequest.AddressID,
	}
}
