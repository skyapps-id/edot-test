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

	public := server.Group("")
	public.Use(JWTMiddleware(container.Config.JwtSecret))
	{
		public.POST("/orders", order.Create)
		public.GET("/orders", order.Gets)
		public.GET("/orders/:uuid", order.Get)
		public.PUT("/orders/:uuid/payment", order.UpdateStatusToPayment)
	}
}
