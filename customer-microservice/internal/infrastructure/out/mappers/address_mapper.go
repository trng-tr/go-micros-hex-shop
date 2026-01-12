package mappers

import (
	"database/sql"

	"github.com/trng-tr/customer-microservice/internal/domain"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/models"
)

func ToAddressModel(bs domain.BusinessAddress) models.Address {
	var complement sql.NullString
	if bs.Complement != nil {
		complement = sql.NullString{
			String: *bs.Complement,
			Valid:  true,
		}
	}
	var streetNumber sql.NullString

	if bs.StreetNumber != nil {
		streetNumber = sql.NullString{
			String: *bs.StreetNumber,
			Valid:  true,
		}
	}

	return models.Address{
		ID:           bs.ID,
		StreetNumber: streetNumber,
		StreetName:   bs.StreetName,
		ZipCode:      bs.ZipCode,
		City:         bs.City,
		Region:       bs.Region,
		Country:      bs.Country,
		Complement:   complement,
	}
}

func ToBusinessAddress(m models.Address) domain.BusinessAddress {
	var complement *string
	if m.Complement.Valid == true {
		complement = &m.Complement.String
	} else {
		complement = nil
	}
	var streetNumber *string
	if m.StreetNumber.Valid {
		streetNumber = &m.StreetNumber.String
	}
	return domain.BusinessAddress{
		ID:           m.ID,
		StreetNumber: streetNumber,
		StreetName:   m.StreetName,
		ZipCode:      m.ZipCode,
		City:         m.City,
		Region:       m.Region,
		Country:      m.Country,
		Complement:   complement,
	}
}
