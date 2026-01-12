package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/models"
)

/*
AddressRepositoryImpl implement AddressRepository,
it uses by DI sql.DB
*/
type AddressRepositoryImpl struct {
	db *sql.DB
}

// NewAddressRepositoryImpl function constructor
func NewAddressRepositoryImpl(db *sql.DB) *AddressRepositoryImpl {
	return &AddressRepositoryImpl{db: db}
}

// SaveO implement interface AddressRepository
func (ari *AddressRepositoryImpl) SaveO(ctx context.Context, o models.Address) (models.Address, error) {
	var query string = `INSERT INTO addresses (street_number,street_name,zip_code,city,region,country,complement)
	VALUES(
	$1,$2,$3,$4,$5,$6,$7)
	RETURNING id`

	var err error = ari.db.QueryRowContext(
		ctx,
		query,
		o.StreetNumber,
		o.StreetName,
		o.ZipCode,
		o.City,
		o.Region,
		o.Country,
		o.Complement,
	).Scan(&o.ID)
	if err != nil {
		return models.Address{}, fmt.Errorf("an error occurred %w", err)
	}

	return o, nil
}

// FindOByID implement interface AddressRepository
func (ari *AddressRepositoryImpl) FindOByID(ctx context.Context, id int64) (models.Address, error) {
	var query string = `SELECT id,street_number,street_name,zip_code,city,region,country,complement
	FROM addresses
	WHERE id=$1`
	var address models.Address
	var err error = ari.db.QueryRowContext(ctx, query, id).Scan(
		&address.ID,
		&address.StreetNumber,
		&address.StreetName,
		&address.ZipCode,
		&address.City,
		&address.Region,
		&address.Country,
		&address.Complement,
	)
	if err != nil {
		return models.Address{}, fmt.Errorf("an error occurred %w", err)
	}
	return address, nil
}

// FindAllO implement interface AddressRepository
func (ari *AddressRepositoryImpl) FindAllO(ctx context.Context) ([]models.Address, error) {
	var query string = `SELECT id,street_number,street_name,zip_code,city,region,country,complement
	FROM addresses`
	rows, err := ari.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var addresses []models.Address = make([]models.Address, 0)
	for rows.Next() {
		var address models.Address
		rows.Scan(
			&address.ID,
			&address.StreetNumber,
			&address.StreetName,
			&address.ZipCode,
			&address.City,
			&address.Region,
			&address.Country,
			&address.Complement,
		)
		addresses = append(addresses, address)
	}

	return addresses, nil
}

// DeleteO implement interface AddressRepository
func (ari *AddressRepositoryImpl) DeleteO(ctx context.Context, id int64) error {
	var query string = "DELETE FROM addresses WHERE id=$1"
	result, err := ari.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
