package order

import (
	"context"

	"github.com/skyapps-id/edot-test/order-service/config"
	"github.com/skyapps-id/edot-test/order-service/repository"
	"github.com/skyapps-id/edot-test/order-service/wrapper/product_service"
)

type OrderUsecase interface {
	Craete(ctx context.Context, req CreateOrderRequest) (resp CreateOrderResponse, err error)
	Gets(ctx context.Context, req GetOrdersRequest) (resp GetOrdersResponse, err error)
	Get(ctx context.Context, req GetOrderRequest) (resp GetOrderResponse, err error)
}

type usecase struct {
	cfg             config.Config
	orderRepository repository.Order
	productWrapper  product_service.ProductServiceWrapper
}

func NewUsecase(cfg config.Config, orderRepository repository.Order, productWrapper product_service.ProductServiceWrapper) OrderUsecase {
	return &usecase{
		cfg:             cfg,
		orderRepository: orderRepository,
		productWrapper:  productWrapper,
	}
}
