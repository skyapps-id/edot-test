package warehouse

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/response"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/validator"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/usecase/warehouse"
)

func (h *handler) GetMaxQuantityByProductUUIDs(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req warehouse.GetWarehouseProductRequest

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	resp, err := h.warehouseUsecase.GetMaxQuantityByProductUUIDs(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}

func (h *handler) GetProductStock(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req warehouse.GetProductStockRequest

	err = c.Bind(&req)
	if err != nil {
		return
	}

	resp, err := h.warehouseUsecase.GetProductStock(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
