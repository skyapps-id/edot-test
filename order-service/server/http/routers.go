package http

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/order-service/container"
	"github.com/skyapps-id/edot-test/order-service/handler/health"
	"github.com/skyapps-id/edot-test/order-service/handler/order"
)

func Router(server *echo.Echo, container *container.Container) {
	// Handler
	health := health.NewHandler()
	order := order.NewHandler(container.OrderUsecase)

	server.GET("/", health.Health)
	server.POST("/orders", order.Create)
	server.GET("/orders", order.Gets)
	server.GET("/orders/:uuid", order.Get)
	server.PUT("/orders/:uuid/payment", order.UpdateStatusToPayment)
}
