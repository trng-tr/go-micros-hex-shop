package mappers

import (
	"database/sql"
	"time"

	"github.com/trng-tr/customer-microservice/internal/domain"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/dtos"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/models"
)

func ToCustomerModel(bs domain.BusinessCustomer) models.Customer {
	var updatedAt sql.NullTime
	if bs.UpdatedAt != nil {
		updatedAt = sql.NullTime{
			Time:  *bs.UpdatedAt,
			Valid: true,
		}
	}

	return models.Customer{
		ID:          bs.ID,
		Firstname:   bs.Firstname,
		Lastname:    bs.Lastname,
		Genda:       string(bs.Genda),
		Email:       bs.Email,
		PhoneNumber: bs.PhoneNumber,
		Status:      string(bs.Status),
		AddressID:   bs.AddressID,
		CreatedAt:   bs.CreatedAt,
		UpdatedAt:   updatedAt,
	}
}

func ToBusinessCustomer(m models.Customer) domain.BusinessCustomer {
	var updatedAt *time.Time
	if m.UpdatedAt.Valid == true {
		updatedAt = &m.UpdatedAt.Time
	} else {
		updatedAt = nil
	}

	return domain.BusinessCustomer{
		ID:          m.ID,
		Firstname:   m.Firstname,
		Lastname:    m.Lastname,
		Genda:       domain.Genda(m.Genda),
		Email:       m.Email,
		PhoneNumber: m.PhoneNumber,
		Status:      domain.Status(m.Status),
		AddressID:   m.AddressID,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   updatedAt,
	}
}

func ApplyPatchCustomer(request dtos.CustomerPatchRequest, m *models.Customer) {
	if request.Firstname != nil {
		m.Firstname = *request.Firstname
	}
	if request.Lastname != nil {
		m.Lastname = *request.Lastname
	}

	if request.PhoneNumber != nil {
		m.PhoneNumber = *request.PhoneNumber
	}

	if request.AddressID != nil {
		m.AddressID = *request.AddressID
	}
}
