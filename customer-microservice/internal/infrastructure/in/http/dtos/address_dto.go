package dtos

type AddressRequest struct {
	StreetNumber *string `json:"street_number,omitempty" binding:"omitempty"`
	StreetName   string  `json:"street_name" binding:"required,min=2"`
	ZipCode      string  `json:"zip_code" binding:"required,min=4,max=10"`
	City         string  `json:"city" binding:"required,min=2"`
	Region       string  `json:"region" binding:"required,min=2"`
	Country      string  `json:"country" binding:"required,min=2"`
	Complement   *string `json:"complement,omitempty"`
}

type AddressResponse struct {
	ID           int64   `json:"id"`
	StreetNumber *string `json:"street_number,omitempty"`
	StreetName   string  `json:"street_name"`
	ZipCode      string  `json:"zip_code"`
	City         string  `json:"city"`
	Region       string  `json:"region"`
	Country      string  `json:"country"`
	Complement   *string `json:"complement,omitempty"`
}
