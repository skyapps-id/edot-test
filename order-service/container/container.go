package container

import (
	"github.com/RichardKnop/machinery/v1"
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
	Worker       *machinery.Server
}

func Setup() *Container {
	// Load Config
	config := config.Load()

	// Machinery
	worker := driver.GetMachineryServer(config)

	// Database
	database := driver.NewGormDatabase(config)

	// Repository
	orderRepository := repository.NewOrderRepository(database)
	orderItemRepository := repository.NewOrderItemRepository(database)

	// Wraper
	productWrapper := product_service.NewWrapper(config).Setup()
	shopWarehouseWrapper := shop_warehouse_service.NewWrapper(config).Setup()

	// Usecase
	orderUsecase := order.NewUsecase(config, worker, orderRepository, orderItemRepository, productWrapper, shopWarehouseWrapper)

	return &Container{
		Config:       config,
		Worker:       worker,
		OrderUsecase: orderUsecase,
	}
}
