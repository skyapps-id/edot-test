package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/product-service/config"
	"github.com/skyapps-id/edot-test/product-service/repository"
	"github.com/skyapps-id/edot-test/product-service/wrapper/shop_warehouse_service"
)

type ProductUsecase interface {
	Craete(ctx context.Context, req CreateProductRequest) (resp CreateProductResponse, err error)
	Gets(ctx context.Context, req GetProductsRequest) (resp GetProductsResponse, err error)
	Get(ctx context.Context, req GetProductRequest) (resp GetProductResponse, err error)
	GetByUUIDs(ctx context.Context, req GetProductByUUIDsRequest) (resp map[uuid.UUID]GetProductResponse, err error)
}

type usecase struct {
	cfg                  config.Config
	productRepository    repository.Product
	shopWarehouseWrapper shop_warehouse_service.ShopWarehousetServiceWrapper
}

func NewUsecase(cfg config.Config, productRepository repository.Product, shopWarehouseWrapper shop_warehouse_service.ShopWarehousetServiceWrapper) ProductUsecase {
	return &usecase{
		cfg:                  cfg,
		productRepository:    productRepository,
		shopWarehouseWrapper: shopWarehouseWrapper,
	}
}
