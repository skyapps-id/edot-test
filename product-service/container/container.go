package container

import (
	"github.com/skyapps-id/edot-test/product-service/config"
	"github.com/skyapps-id/edot-test/product-service/driver"
	"github.com/skyapps-id/edot-test/product-service/repository"
	"github.com/skyapps-id/edot-test/product-service/usecase/product"
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
	repo_user := repository.NewProductRepository(database)

	// Usecase
	productUsecase := product.NewUsecase(config, repo_user)

	return &Container{
		Config:         config,
		ProductUsecase: productUsecase,
	}
}
