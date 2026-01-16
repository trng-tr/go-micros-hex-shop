package mappers

import (
	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
)

// ToStockModel mapper to object for db
func ToStockModel(bsStock domain.Stock) models.StockModel {

	return models.StockModel{
		ID:        bsStock.ID,
		ProductID: bsStock.ProductID,
		Quantity:  bsStock.Quantity,
		UpdatedAt: bsStock.UpdatedAt,
	}
}

func ToBusinessStock(model models.StockModel) domain.Stock {
	return domain.Stock{
		ID:        model.ID,
		ProductID: model.ProductID,
		Quantity:  model.Quantity,
		UpdatedAt: model.UpdatedAt,
	}
}
