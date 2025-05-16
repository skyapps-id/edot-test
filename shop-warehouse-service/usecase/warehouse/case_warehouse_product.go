package warehouse

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/tracer"
	"gorm.io/gorm"
)

func (uc *usecase) CreateWarehouseProduct(ctx context.Context, req CreateWarehouseProductRequest) (resp CreateWarehouseProductResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.Create")
	defer span.End()

	_, err = uc.warehouseRepository.FindByUUID(ctx, req.WarehouseUUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("shop not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get warehouse"))
		return
	}

	// Validation Product UUID

	err = uc.warehouseProductRepository.Create(ctx, entity.WarehouseProduct{
		WarehouseUUID: req.WarehouseUUID,
		ProductUUID:   req.ProductUUID,
		Quantity:      req.Quantity,
	})
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to save warehouse product"))
		return
	}

	resp.WarehouseUUID = req.WarehouseUUID
	resp.ProductUUID = req.ProductUUID
	resp.Quantity = req.Quantity

	return
}
