package warehouse

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/response"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/validator"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/usecase/warehouse"
)

func (h *handler) WarehouseUpdateActive(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req warehouse.WarehouseUpdateActiveRequest

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	resp, err := h.warehouseUsecase.WarehouseUpdateActive(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
