package product

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/product-service/usecase/product"
)

type ProductHandler interface {
	Create(c echo.Context) (err error)
	Gets(c echo.Context) (err error)
}

type handler struct {
	productUsecase product.ProductUsecase
}

func NewHandler(usecase product.ProductUsecase) ProductHandler {
	return &handler{
		productUsecase: usecase,
	}
}
