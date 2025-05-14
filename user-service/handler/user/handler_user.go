package user

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/user-service/pkg/response"
	"github.com/skyapps-id/edot-test/user-service/usecase/user"
)

func (h *handler) Register(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req user.RegisterRequest

	err = c.Bind(&req)
	if err != nil {
		return
	}

	resp, err := h.userUsecase.Register(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
