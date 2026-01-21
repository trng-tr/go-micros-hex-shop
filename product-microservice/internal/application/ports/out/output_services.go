package out

import (
	"context"

	"github.com/trng-tr/product-microservice/internal/domain"
)

// OutProductService contrat pour la gestion du stock
type OutProductService interface {
	SaveProduct(ctx context.Context, prd domain.Product) (domain.Product, error)
	GetProductByID(ctx context.Context, id int64) (domain.Product, error)
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
	PatchProduct(ctx context.Context, id int64, product domain.Product) (domain.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
	GetProductBySku(ctx context.Context, sku string) (domain.Product, error)
}

// OutStockService contrat pour la gestion du stock
type OutStockService interface {
	CreateStock(ctx context.Context, stk domain.Stock) (domain.Stock, error)
	GetStockByID(ctx context.Context, id int64) (domain.Stock, error)
	GetAllStocks(ctx context.Context) ([]domain.Stock, error)
	UpdateStockQuantity(ctx context.Context, stock domain.Stock) (domain.Stock, error)
	GetStockByProductID(ctx context.Context, productID int64) (domain.Stock, error)
}

// OutGenerateSkuService service pour générer un sku du produit
type OutUuidGeneratorService interface {
	GenerateUuid() string
}
