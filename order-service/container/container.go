package container

import (
	"github.com/skyapps-id/edot-test/order-service/config"
	"github.com/skyapps-id/edot-test/order-service/driver"
	"github.com/skyapps-id/edot-test/order-service/repository"
	"github.com/skyapps-id/edot-test/order-service/usecase/order"
	"github.com/skyapps-id/edot-test/order-service/wrapper/product_service"
	"github.com/skyapps-id/edot-test/order-service/wrapper/shop_warehouse_service"
)

type Container struct {
	Config       config.Config
	OrderUsecase order.OrderUsecase
}

func Setup() *Container {
	// Load Config
	config := config.Load()

	// Database
	database := driver.NewGormDatabase(config)

	// Repository
	repo_order := repository.NewOrderRepository(database)

	// Wraper
	product_wrapper := product_service.NewWrapper(config).Setup()
	shop_warehouse_wrapper := shop_warehouse_service.NewWrapper(config).Setup()

	// Usecase
	orderUsecase := order.NewUsecase(config, repo_order, product_wrapper, shop_warehouse_wrapper)

	return &Container{
		Config:       config,
		OrderUsecase: orderUsecase,
	}
}
