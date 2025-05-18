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

	public := server.Group("")
	public.Use(JWTMiddleware(container.Config.JwtSecret))
	{
		public.POST("/shops", shop.Create)
		public.GET("/shops", shop.Gets)
		public.GET("/shops/:uuid", shop.Get)
		public.POST("/warehouses", warehouse.Create)
		public.GET("/warehouses", warehouse.Gets)
		public.GET("/warehouses/:uuid", warehouse.Get)
		public.PUT("/warehouses/:uuid/:status", warehouse.WarehouseUpdateActive)
		public.POST("/warehouses/products", warehouse.CreateWarehouseProduct)
		public.PUT("/warehouses/product-restock", warehouse.ProductRestock)
		public.PUT("/warehouses/product-transfer-stock", warehouse.TransferStock)
	}

	internal := server.Group("/internal")
	internal.Use(ValidateStaticToken(container.Config.TokenInternal))
	{
		internal.GET("/warehouses/product-stock/:uuid", warehouse.GetProductStock)
		internal.POST("/warehouses/product-stock", warehouse.GetMaxQuantityByProductUUIDs)
		internal.POST("/warehouses/product-stock-addition", warehouse.ProductStockAddition)
		internal.POST("/warehouses/product-stock-reduction", warehouse.ProductStockReduction)
	}

}
