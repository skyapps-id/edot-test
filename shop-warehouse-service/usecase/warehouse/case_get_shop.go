package warehouse

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/tracer"
	util "github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/utils"
	"gorm.io/gorm"
)

func (uc *usecase) Gets(ctx context.Context, req GetWarehousesRequest) (resp GetWarehousesResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.Gets")
	defer span.End()

	warehouses, count, err := uc.warehouseRepository.GetAll(
		ctx,
		req.Name,
		int(req.PerPage.Int64), int(req.Page.Int64), req.Sort)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to gets warehouse"))
		return
	}

	resp.List = uc.warehousesMapper(warehouses)
	resp.Pagination = util.Pagination(int(req.Page.Int64), int(req.PerPage.Int64), count)

	return
}

func (uc *usecase) Get(ctx context.Context, req GetWarehouseRequest) (resp GetWarehouseResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.Get")
	defer span.End()

	warehouse, err := uc.warehouseRepository.FindByUUID(ctx, req.UUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("shop not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get warehouse"))
		return
	}

	resp = uc.warehouseMapper(warehouse)

	return
}
