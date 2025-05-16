package warehouse

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/tracer"
	"gorm.io/gorm"
)

func (uc *usecase) GetMaxQuantityByProductUUIDs(ctx context.Context, req GetWarehouseProductRequest) (resp map[uuid.UUID]GetWarehouseProductResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.GetWarehouseProduct")
	defer span.End()

	warehouseProduct, err := uc.warehouseProductRepository.GetMaxQuantityByProductUUIDs(ctx, req.ProductUUIDs)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("shop not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get warehouse product"))
		return
	}

	resp = uc.warehouseProductMapper(req, warehouseProduct)

	return
}
