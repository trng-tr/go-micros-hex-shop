package models

import "database/sql"

//Address model of data for table addresses
type Address struct {
	ID           int64
	StreetNumber sql.NullString //because it is not mandatory
	StreetName   string
	ZipCode      string
	City         string
	Region       string
	Country      string
	Complement   sql.NullString //because it is not mandatory
}
