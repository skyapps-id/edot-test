package container

import (
	"github.com/skyapps-id/edot-test/order-service/config"
	"github.com/skyapps-id/edot-test/order-service/driver"
	"github.com/skyapps-id/edot-test/order-service/repository"
	"github.com/skyapps-id/edot-test/order-service/usecase/order"
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

	// Usecase
	orderUsecase := order.NewUsecase(config, repo_order)

	return &Container{
		Config:       config,
		OrderUsecase: orderUsecase,
	}
}
