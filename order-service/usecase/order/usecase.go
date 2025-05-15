package order

import (
	"context"

	"github.com/skyapps-id/edot-test/order-service/config"
	"github.com/skyapps-id/edot-test/order-service/repository"
)

type OrderUsecase interface {
	Craete(ctx context.Context, req CreateOrderRequest) (resp CreateOrderResponse, err error)
	Gets(ctx context.Context, req GetOrdersRequest) (resp GetOrdersResponse, err error)
	Get(ctx context.Context, req GetOrderRequest) (resp GetOrderResponse, err error)
}

type usecase struct {
	cfg             config.Config
	orderRepository repository.Order
}

func NewUsecase(cfg config.Config, orderRepository repository.Order) OrderUsecase {
	return &usecase{
		cfg:             cfg,
		orderRepository: orderRepository,
	}
}
