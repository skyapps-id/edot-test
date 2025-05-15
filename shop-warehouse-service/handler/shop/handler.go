package shop

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/usecase/shop"
)

type ShopHandler interface {
	Create(c echo.Context) (err error)
	Gets(c echo.Context) (err error)
	Get(c echo.Context) (err error)
}

type handler struct {
	shopUsecase shop.ShopUsecase
}

func NewHandler(usecase shop.ShopUsecase) ShopHandler {
	return &handler{
		shopUsecase: usecase,
	}
}
