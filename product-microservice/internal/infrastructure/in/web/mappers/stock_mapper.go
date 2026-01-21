package mappers

import (
	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/dtos"
)

func ToStockResponse(bsStock domain.Stock, bsProduct domain.Product) dtos.StockResponse {
	return dtos.StockResponse{
		ID:              bsStock.ID,
		Name:            bsStock.Name,
		ProductId:       bsStock.ProductID,
		Quantity:        bsStock.Quantity,
		UpdatedAt:       bsStock.UpdatedAt.Format(dateFormate),
		ProductResponse: ToProductResponse(bsProduct),
	}
}

func ToBusinessStock(request dtos.StockRequest) domain.Stock {
	return domain.Stock{
		Name:      request.Name,
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}
}
