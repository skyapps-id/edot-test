package http

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/container"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/handler/health"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/handler/shop"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/handler/warehouse"
)

func Router(server *echo.Echo, container *container.Container) {
	// Handler
	health := health.NewHandler()
	shop := shop.NewHandler(container.ShopUsecase)
	warehouse := warehouse.NewHandler(container.WarehouseUsecase)

	server.GET("/", health.Health)
	server.POST("/shops", shop.Create)
	server.GET("/shops", shop.Gets)
	server.GET("/shops/:uuid", shop.Get)
	server.POST("/warehouses", warehouse.Create)
	server.GET("/warehouses", warehouse.Gets)
	server.GET("/warehouses/:uuid", warehouse.Get)
	server.POST("/warehouses/products", warehouse.CreateWarehouseProduct)
}
