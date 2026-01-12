package services

import (
	"context"
	"fmt"

	"github.com/trng-tr/customer-microservice/internal/domain"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/mappers"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/models"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/repositories/contract"
)

/*
OutAddressServiceImpl implementation of output ports,
to persit in db, we inject address repository
*/
type OutAddressServiceImpl struct {
	repo contract.AddressRepository
}

// NewOutAddressServiceImpl func constructor
func NewOutAddressServiceImpl(repo contract.AddressRepository) *OutAddressServiceImpl {
	return &OutAddressServiceImpl{repo: repo}
}

// CreateAddress implement OutAddressService : output port
func (oasi *OutAddressServiceImpl) CreateAddress(ctx context.Context, bsAddress domain.BusinessAddress) (
	domain.BusinessAddress, error) {
	var addressModel models.Address = mappers.ToAddressModel(bsAddress)
	savedAddress, err := oasi.repo.SaveO(ctx, addressModel)
	if err != nil {
		return domain.BusinessAddress{}, fmt.Errorf("create address faild %w", err) //%w pour wrapper err
	}
	return mappers.ToBusinessAddress(savedAddress), nil
}

// GetAddressByID implement OutAddressService : output port
func (oasi *OutAddressServiceImpl) GetAddressByID(ctx context.Context, id int64) (domain.BusinessAddress, error) {
	addressModel, err := oasi.repo.FindOByID(ctx, id)
	if err != nil {
		return domain.BusinessAddress{}, fmt.Errorf("find address by id faild %w", err) //%w pour wrapper err
	}

	return mappers.ToBusinessAddress(addressModel), nil
}

// GetAllAddresses implement OutAddressService : output port
func (oasi *OutAddressServiceImpl) GetAllAddresses(ctx context.Context) ([]domain.BusinessAddress, error) {
	addressesModels, err := oasi.repo.FindAllO(ctx)
	if err != nil {
		return nil, err
	}
	var bsAddresses []domain.BusinessAddress = make([]domain.BusinessAddress, 0, len(addressesModels))
	for _, aModel := range addressesModels {
		var bsAddress = mappers.ToBusinessAddress(aModel)
		bsAddresses = append(bsAddresses, bsAddress)
	}
	return bsAddresses, nil
}

// DeleteAddress implement OutAddressService : output port
func (oasi *OutAddressServiceImpl) DeleteAddress(ctx context.Context, id int64) error {
	if err := oasi.repo.DeleteO(ctx, id); err != nil {
		return err
	}

	return nil
}
