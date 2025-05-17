package warehouse

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/tracer"
)

func (uc *usecase) ProductStockAddition(ctx context.Context, req ProductStockAdditionRequest) (resp ProductStockAdditionResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.ProductStockAddition")
	defer span.End()

	productUUIDs := make([]uuid.UUID, 0, len(req.Products))
	products := make([]entity.ProductStock, 0, len(req.Products))

	for _, row := range req.Products {
		productUUIDs = append(productUUIDs, row.ProductUUID)
		products = append(products, entity.ProductStock{
			ProductUUID:   row.ProductUUID,
			WarehouseUUID: row.WarehouseUUID,
			Quantity:      row.Quantity,
		})
	}

	err = uc.warehouseProductRepository.ProductStockAddition(ctx, products)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to product stock addition"))
		return
	}

	resp.ProductUUIDs = productUUIDs

	return
}

func (uc *usecase) ProductStockReduction(ctx context.Context, req ProductStockReductionRequest) (resp ProductStockReductionResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.ProductStockReduction")
	defer span.End()

	productUUIDs := make([]uuid.UUID, 0, len(req.Products))
	products := make([]entity.ProductStock, 0, len(req.Products))

	for _, row := range req.Products {
		productUUIDs = append(productUUIDs, row.ProductUUID)
		products = append(products, entity.ProductStock{
			ProductUUID:   row.ProductUUID,
			WarehouseUUID: row.WarehouseUUID,
			Quantity:      row.Quantity,
		})
	}

	err = uc.warehouseProductRepository.ProductStockReduction(ctx, products)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to product stock reduction"))
		return
	}

	resp.ProductUUIDs = productUUIDs

	return
}
