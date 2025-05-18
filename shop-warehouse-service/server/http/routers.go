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
	server.PUT("/warehouses/:uuid/:status", warehouse.WarehouseUpdateActive)
	server.POST("/warehouses/products", warehouse.CreateWarehouseProduct)
	server.PUT("/warehouses/product-restock", warehouse.ProductRestock)
	server.PUT("/warehouses/product-transfer-stock", warehouse.TransferStock)
	internal := server.Group("/internal")
	internal.Use(ValidateStaticToken(container.Config.TokenInternal))
	{
		internal.GET("/warehouses/product-stock/:uuid", warehouse.GetProductStock)
		internal.POST("/warehouses/product-stock", warehouse.GetMaxQuantityByProductUUIDs)
		internal.POST("/warehouses/product-stock-addition", warehouse.ProductStockAddition)
		internal.POST("/warehouses/product-stock-reduction", warehouse.ProductStockReduction)
	}

}
