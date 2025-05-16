package warehouse

import (
	"context"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/tracer"
)

func (uc *usecase) Craete(ctx context.Context, req CreateWarehouseRequest) (resp CreateWarehouseResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.Create")
	defer span.End()

	err = uc.warehouseRepository.Create(ctx, entity.Warehouse{
		Name:     req.Name,
		Address:  req.Address,
		ShopUUID: req.ShopUUID,
		Active:   true,
	})
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to save warehouse"))
		return
	}

	resp.Name = req.Name

	return
}
