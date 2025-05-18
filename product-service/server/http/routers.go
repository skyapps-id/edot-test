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

	public := server.Group("")
	public.Use(JWTMiddleware(container.Config.JwtSecret))
	{
		public.POST("/products", product.Create)
		public.GET("/products", product.Gets)
		public.GET("/products/:uuid", product.Get)
	}

	internal := server.Group("/internal")
	internal.Use(ValidateStaticToken(container.Config.TokenInternal))
	{
		internal.POST("/products/uuids", product.GetByUUIDs)
	}
}
