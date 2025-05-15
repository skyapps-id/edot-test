package health

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/product-service/pkg/logger"
	"github.com/skyapps-id/edot-test/product-service/pkg/response"
)

type HealthHandler interface {
	Health(c echo.Context) (err error)
}

type handler struct {
}

func NewHandler() HealthHandler {
	return &handler{}
}

func (h *handler) Health(c echo.Context) (err error) {

	logger.Log.Info("Health")
	return response.ResponseSuccess(c, "Ok")
}
