package mappers

import (
	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/dtos"
)

func ToStockResponse(bsStock domain.Stock, bsProduct domain.Product) dtos.StockResponse {
	return dtos.StockResponse{
		ID:              bsStock.ID,
		ProductId:       bsStock.ProductID,
		Quantity:        bsStock.Quantity,
		UpdatedAt:       bsStock.UpdatedAt,
		ProductResponse: ToProductResponse(bsProduct),
	}
}

func ToBusinessStock(request dtos.StockRequest) domain.Stock {
	return domain.Stock{
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}
}
