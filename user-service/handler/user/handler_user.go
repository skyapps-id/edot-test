package user

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/user-service/pkg/response"
	"github.com/skyapps-id/edot-test/user-service/pkg/validator"
	"github.com/skyapps-id/edot-test/user-service/usecase/user"
)

func (h *handler) Register(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req user.RegisterRequest

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	resp, err := h.userUsecase.Register(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}

func (h *handler) Login(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req user.LoginRequest

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	resp, err := h.userUsecase.Login(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
