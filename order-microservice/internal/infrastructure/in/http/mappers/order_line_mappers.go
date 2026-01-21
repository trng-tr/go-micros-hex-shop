package mappers

import (
	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
)

func ToBusinessOrderLine(request dtos.OrderLineRequest) domain.OrderLine {
	return domain.OrderLine{
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}
}

func ToOrderLineResponse(orderLine domain.OrderLine, product domain.Product) dtos.OrderLineResponse {
	return dtos.OrderLineResponse{
		ID: orderLine.ID,
		ProductResponse: dtos.ProductResponse{
			ID:          product.ID,
			Sku:         product.Sku,
			ProductName: product.ProductName,
			Description: product.Description,
			PriceResponse: dtos.PriceResponse{
				UnitPrice: product.Price.UnitPrice,
				Currency:  string(product.Price.Currency),
			},
			IsActive: product.IsActive,
		},
		Quantity: orderLine.Quantity,
	}
}
