package order

import (
	"context"

	"github.com/RichardKnop/machinery/v1"
	"github.com/skyapps-id/edot-test/order-service/config"
	"github.com/skyapps-id/edot-test/order-service/repository"
	"github.com/skyapps-id/edot-test/order-service/wrapper/product_service"
	"github.com/skyapps-id/edot-test/order-service/wrapper/shop_warehouse_service"
)

type OrderUsecase interface {
	Craete(ctx context.Context, req CreateOrderRequest) (resp CreateOrderResponse, err error)
	Gets(ctx context.Context, req GetOrdersRequest) (resp GetOrdersResponse, err error)
	Get(ctx context.Context, req GetOrderRequest) (resp GetOrderResponse, err error)
	UpdateStatusToPayment(ctx context.Context, req OrderStatusToPaymentRequest) (resp OrderStatusToPaymentResponse, err error)
	OrderCancel(ctx context.Context, req OrderCancelRequest) (resp OrderCancelResponse, err error)
}

type usecase struct {
	cfg                  config.Config
	worker               *machinery.Server
	orderRepository      repository.Order
	orderItemRepository  repository.OrderItem
	productWrapper       product_service.ProductServiceWrapper
	shopWarehouseWrapper shop_warehouse_service.ShopWarehousetServiceWrapper
}

func NewUsecase(
	cfg config.Config,
	worker *machinery.Server,
	orderRepository repository.Order,
	orderItemRepository repository.OrderItem,
	productWrapper product_service.ProductServiceWrapper,
	shopWarehouseWrapper shop_warehouse_service.ShopWarehousetServiceWrapper,
) OrderUsecase {
	return &usecase{
		cfg:                  cfg,
		worker:               worker,
		orderRepository:      orderRepository,
		orderItemRepository:  orderItemRepository,
		productWrapper:       productWrapper,
		shopWarehouseWrapper: shopWarehouseWrapper,
	}
}
