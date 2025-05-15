package http

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/product-service/container"
	"github.com/skyapps-id/edot-test/product-service/handler/health"
	"github.com/skyapps-id/edot-test/product-service/handler/product"
)

func Router(server *echo.Echo, container *container.Container) {
	// Handler
	health := health.NewHandler()
	product := product.NewHandler(container.ProductUsecase)

	server.GET("/", health.Health)
	server.POST("/products", product.Create)
}
