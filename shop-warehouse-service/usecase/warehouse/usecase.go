package warehouse

import (
	"context"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/config"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/repository"
)

type WarehouseUsecase interface {
	Craete(ctx context.Context, req CreateWarehouseRequest) (resp CreateWarehouseResponse, err error)
	Gets(ctx context.Context, req GetWarehousesRequest) (resp GetWarehousesResponse, err error)
	Get(ctx context.Context, req GetWarehouseRequest) (resp GetWarehouseResponse, err error)
	CreateWarehouseProduct(ctx context.Context, req CreateWarehouseProductRequest) (resp CreateWarehouseProductResponse, err error)
	GetMaxQuantityByProductUUIDs(ctx context.Context, req GetWarehouseProductRequest) (resp map[uuid.UUID]GetWarehouseProductResponse, err error)
	ProductStockAddition(ctx context.Context, req ProductStockAdditionRequest) (resp ProductStockAdditionResponse, err error)
	ProductStockReduction(ctx context.Context, req ProductStockReductionRequest) (resp ProductStockReductionResponse, err error)
	GetProductStock(ctx context.Context, req GetProductStockRequest) (resp GetProductStockResponse, err error)
	WarehouseUpdateActive(ctx context.Context, req WarehouseUpdateActiveRequest) (resp WarehouseUpdateActiveResponse, err error)
	ProductRestock(ctx context.Context, req ProductRestockReqeust) (resp ProductRestockResponse, err error)
	TransferStock(ctx context.Context, req TransferStockReqeust) (resp TransferStockResponse, err error)
}

type usecase struct {
	cfg                        config.Config
	warehouseRepository        repository.Warehouse
	warehouseProductRepository repository.WarehouseProduct
}

func NewUsecase(cfg config.Config, warehouseRepository repository.Warehouse, warehouseProductRepository repository.WarehouseProduct) WarehouseUsecase {
	return &usecase{
		cfg:                        cfg,
		warehouseRepository:        warehouseRepository,
		warehouseProductRepository: warehouseProductRepository,
	}
}
