package user

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/user-service/usecase/user"
)

type UserHandler interface {
	Register(c echo.Context) (err error)
	Login(c echo.Context) (err error)
}

type handler struct {
	userUsecase user.UserUsecase
}

func NewHandler(usecase user.UserUsecase) UserHandler {
	return &handler{
		userUsecase: usecase,
	}
}
