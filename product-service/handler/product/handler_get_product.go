package product

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/product-service/pkg/response"
	"github.com/skyapps-id/edot-test/product-service/usecase/product"
	"gopkg.in/guregu/null.v4"
)

func (h *handler) Gets(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req product.GetProductsRequest

	err = c.Bind(&req)
	if err != nil {
		return
	}

	if req.Sort == "" {
		req.Sort = "DESC"
	}
	req.Sort = strings.ToUpper(req.Sort)

	if !req.PerPage.Valid {
		req.PerPage = null.NewInt(10, true)
	}

	if !req.Page.Valid {
		req.Page = null.NewInt(1, true)
	}

	resp, err := h.productUsecase.Gets(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}

func (h *handler) Get(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req product.GetProductRequest

	err = c.Bind(&req)
	if err != nil {
		return
	}

	resp, err := h.productUsecase.Get(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
