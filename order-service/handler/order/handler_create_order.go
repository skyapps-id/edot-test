package order

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/auth"
	"github.com/skyapps-id/edot-test/order-service/pkg/response"
	"github.com/skyapps-id/edot-test/order-service/pkg/validator"
	"github.com/skyapps-id/edot-test/order-service/usecase/order"
)

func (h *handler) Create(c echo.Context) (err error) {
	var req order.CreateOrderRequest

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	UserID, ok := auth.GetUserID(c)
	if !ok {
		err = apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("fail to get user id"))
		return
	}
	req.UserUUID = UserID

	resp, err := h.orderUsecase.Craete(c.Request().Context(), req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
