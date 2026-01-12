package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/trng-tr/customer-microservice/internal/application/ports/out"
	"github.com/trng-tr/customer-microservice/internal/domain"
	"github.com/trng-tr/customer-microservice/internal/domain/validators"
)

/*
AddressServiceImpl implement port d'entrée exposé à l'extreieur
il utilise pour cela le port de sortie OutAddressService
*/
type InAddressServiceImpl struct {
	outService out.OutAddressService
}

// NewAddressServiceImpl DI by constructor
func NewInAddressServiceImpl(outService out.OutAddressService) *InAddressServiceImpl {
	return &InAddressServiceImpl{outService: outService}
}

// CreateAddress implement InAddressService interface
func (iasi *InAddressServiceImpl) CreateAddress(ctx context.Context, bsAddress domain.BusinessAddress) (domain.BusinessAddress, error) {
	var inputFields map[string]string = map[string]string{
		"street_name": bsAddress.StreetName,
		"zip_code":    bsAddress.ZipCode,
		"city":        bsAddress.City,
		"region":      bsAddress.Region,
		"country":     bsAddress.Country,
	}
	var err error = validators.CheckInputFields(inputFields)
	if err != nil {
		return domain.BusinessAddress{}, err
	}
	// send business object to outside using output port
	sendBsAddress, err := iasi.outService.CreateAddress(ctx, bsAddress)
	if err != nil {
		return domain.BusinessAddress{}, fmt.Errorf("%w", errSendObject.Error())
	}

	return sendBsAddress, nil
}

// GetAddressByID implement InAddressService interface
func (iasi *InAddressServiceImpl) GetAddressByID(ctx context.Context, id int64) (domain.BusinessAddress, error) {
	if err := validators.CheckInputId(id); err != nil {
		return domain.BusinessAddress{}, err
	}
	//call output port to retried business address object
	bsAddress, err := iasi.outService.GetAddressByID(ctx, id)
	if err != nil {
		return domain.BusinessAddress{}, fmt.Errorf("%w", errRetrieveObject)
	}

	return bsAddress, nil
}

// GetAllAddresses implement InAddressService interface
func (iasi *InAddressServiceImpl) GetAllAddresses(ctx context.Context) ([]domain.BusinessAddress, error) {
	addresses, err := iasi.outService.GetAllAddresses(ctx)
	if err != nil {
		return nil, fmt.Errorf("an arror has occurred %w", err)
	}
	if len(addresses) == 0 {
		return nil, errors.New("no address is registered")
	}

	return addresses, nil
}

// DeleteAddress implement InAddressService interface
func (iasi *InAddressServiceImpl) DeleteAddress(ctx context.Context, id int64) error {
	if err := validators.CheckInputId(id); err != nil {
		return fmt.Errorf("error: delete address failed %w", err)
	}
	if err := iasi.outService.DeleteAddress(ctx, id); err != nil {
		return fmt.Errorf("error: delete address failed %w", err)
	}

	return nil
}
