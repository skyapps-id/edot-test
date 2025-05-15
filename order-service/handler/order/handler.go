package order

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/order-service/usecase/order"
)

type OrderHandler interface {
	Create(c echo.Context) (err error)
	Gets(c echo.Context) (err error)
	Get(c echo.Context) (err error)
}

type handler struct {
	orderUsecase order.OrderUsecase
}

func NewHandler(usecase order.OrderUsecase) OrderHandler {
	return &handler{
		orderUsecase: usecase,
	}
}
