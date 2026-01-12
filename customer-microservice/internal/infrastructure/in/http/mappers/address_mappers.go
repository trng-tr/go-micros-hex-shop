package mappers

import (
	"github.com/trng-tr/customer-microservice/internal/domain"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/dtos"
)

func ToBusinessAddress(request dtos.AddressRequest) domain.BusinessAddress {
	return domain.BusinessAddress{
		StreetNumber: request.StreetNumber,
		StreetName:   request.StreetName,
		ZipCode:      request.ZipCode,
		City:         request.City,
		Region:       request.Region,
		Country:      request.Country,
		Complement:   request.Complement,
	}
}

func ToAddressResponse(bs domain.BusinessAddress) dtos.AddressResponse {
	return dtos.AddressResponse{
		ID:           bs.ID,
		StreetNumber: bs.StreetNumber,
		StreetName:   bs.StreetName,
		ZipCode:      bs.ZipCode,
		City:         bs.City,
		Region:       bs.Region,
		Country:      bs.Country,
		Complement:   bs.Complement,
	}
}
