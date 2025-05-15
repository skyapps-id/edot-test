package product

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/product-service/pkg/response"
	"github.com/skyapps-id/edot-test/product-service/pkg/validator"
	"github.com/skyapps-id/edot-test/product-service/usecase/product"
)

func (h *handler) Create(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req product.CreateProductRequest

	err = validator.Validate(c, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	resp, err := h.productUsecase.Craete(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
