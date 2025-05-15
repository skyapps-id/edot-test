package product

import (
	"context"

	"github.com/skyapps-id/edot-test/product-service/config"
	"github.com/skyapps-id/edot-test/product-service/repository"
)

type ProductUsecase interface {
	Craete(ctx context.Context, req CreateProductRequest) (resp CreateProductResponse, err error)
}

type usecase struct {
	cfg            config.Config
	userRepository repository.Product
}

func NewUsecase(cfg config.Config, userRepository repository.Product) ProductUsecase {
	return &usecase{
		cfg:            cfg,
		userRepository: userRepository,
	}
}
