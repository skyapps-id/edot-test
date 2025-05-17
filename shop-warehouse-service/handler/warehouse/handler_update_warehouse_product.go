package warehouse

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/response"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/validator"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/usecase/warehouse"
)

func (h *handler) ProductRestock(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req warehouse.ProductRestockReqeust

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	resp, err := h.warehouseUsecase.ProductRestock(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}

func (h *handler) ProductStockAddition(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req warehouse.ProductStockAdditionRequest

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	resp, err := h.warehouseUsecase.ProductStockAddition(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}

func (h *handler) ProductStockReduction(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req warehouse.ProductStockReductionRequest

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	resp, err := h.warehouseUsecase.ProductStockReduction(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
