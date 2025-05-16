package container

import (
	"github.com/skyapps-id/edot-test/shop-warehouse-service/config"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/driver"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/repository"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/usecase/shop"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/usecase/warehouse"
)

type Container struct {
	Config           config.Config
	ShopUsecase      shop.ShopUsecase
	WarehouseUsecase warehouse.WarehouseUsecase
}

func Setup() *Container {
	// Load Config
	config := config.Load()

	// Database
	database := driver.NewGormDatabase(config)

	// Repository
	repo_shop := repository.NewShopRepository(database)
	repo_warehouse := repository.NewWarehouseRepository(database)
	repo_warehouse_product := repository.NewWarehouseProductRepository(database)

	// Usecase
	shopUsecase := shop.NewUsecase(config, repo_shop)
	warehouseUsecase := warehouse.NewUsecase(config, repo_warehouse, repo_warehouse_product)

	return &Container{
		Config:           config,
		ShopUsecase:      shopUsecase,
		WarehouseUsecase: warehouseUsecase,
	}
}
