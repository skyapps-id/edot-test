package order

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/order-service/pkg/response"
	"github.com/skyapps-id/edot-test/order-service/pkg/validator"
	"github.com/skyapps-id/edot-test/order-service/usecase/order"
)

func (h *handler) Create(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req order.CreateOrderRequest

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	resp, err := h.orderUsecase.Craete(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
