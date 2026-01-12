package dtos

import (
	"time"
)

// CustomerRequest create request
type CustomerRequest struct {
	Firstname   string `json:"firstname" binding:"required,min=2"`
	Lastname    string `json:"lastname" binding:"required,min=2"`
	Genda       string `json:"genda" binding:"required,min=1"`
	Email       string `json:"email" binding:"required,min=5"`
	PhoneNumber string `json:"phone_number" binding:"required,min=8,max=20"`
	AddressID   int64  `json:"address_id" binding:"required"`
}

// CustomerPatchRequest patch request
type CustomerPatchRequest struct {
	Firstname   *string `json:"firstname,omitempty" binding:"omitempty,min=2"`
	Lastname    *string `json:"lastname,omitempty" binding:"omitempty,min=2"`
	Email       *string `json:"email,omitempty" binding:"omitempty,email"`
	PhoneNumber *string `json:"phone_number,omitempty" binding:"omitempty"`
	AddressID   *int64  `json:"address_id,omitempty" binding:"omitempty,gt=0"`
}

type CustomerResponse struct {
	ID              int64           `json:"id"`
	Firstname       string          `json:"firstname"`
	Lastname        string          `json:"lastname"`
	Genda           string          `json:"genda"`
	Email           string          `json:"email"`
	PhoneNumber     string          `json:"phone_number"`
	Status          string          `json:"status"`
	AddressResponse AddressResponse `json:"address"`
	CreatedAt       time.Time       `json:"created_date"`
	UpdatedAt       *time.Time      `json:"updated_date,omitempty"`
}
