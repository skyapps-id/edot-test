package container

import (
	"github.com/skyapps-id/edot-test/product-service/config"
	"github.com/skyapps-id/edot-test/product-service/driver"
	"github.com/skyapps-id/edot-test/product-service/repository"
	"github.com/skyapps-id/edot-test/product-service/usecase/product"
	"github.com/skyapps-id/edot-test/product-service/wrapper/shop_warehouse_service"
)

type Container struct {
	Config         config.Config
	ProductUsecase product.ProductUsecase
}

func Setup() *Container {
	// Load Config
	config := config.Load()

	// Database
	database := driver.NewGormDatabase(config)

	// Repository
	repo_product := repository.NewProductRepository(database)

	// Wrapper
	shop_warehouse_wrapper := shop_warehouse_service.NewWrapper(config).Setup()

	// Usecase
	productUsecase := product.NewUsecase(config, repo_product, shop_warehouse_wrapper)

	return &Container{
		Config:         config,
		ProductUsecase: productUsecase,
	}
}
