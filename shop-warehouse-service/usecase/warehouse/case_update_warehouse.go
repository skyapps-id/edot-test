package warehouse

import (
	"context"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/tracer"
)

func (uc *usecase) WarehouseUpdateActive(ctx context.Context, req WarehouseUpdateActiveRequest) (resp WarehouseUpdateActiveResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.UpdateActive")
	defer span.End()

	isActive := false
	if req.Status == "active" {
		isActive = true
	}

	err = uc.warehouseRepository.WarehouseUpdateActive(ctx, req.UUID, isActive)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to update warehouse status active %w", err))
		return
	}

	resp.UUID = req.UUID

	return
}
