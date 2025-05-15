package shop

import (
	"context"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/config"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/repository"
)

type ShopUsecase interface {
	Craete(ctx context.Context, req CreateShopRequest) (resp CreateShopResponse, err error)
	Gets(ctx context.Context, req GetShopsRequest) (resp GetShopsResponse, err error)
	Get(ctx context.Context, req GetShopRequest) (resp GetShopResponse, err error)
}

type usecase struct {
	cfg            config.Config
	shopRepository repository.Shop
}

func NewUsecase(cfg config.Config, shopRepository repository.Shop) ShopUsecase {
	return &usecase{
		cfg:            cfg,
		shopRepository: shopRepository,
	}
}
