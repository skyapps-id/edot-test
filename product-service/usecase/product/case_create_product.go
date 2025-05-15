package product

import (
	"context"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/product-service/entity"
	"github.com/skyapps-id/edot-test/product-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/product-service/pkg/tracer"
)

func (uc *usecase) Craete(ctx context.Context, req CreateProductRequest) (resp CreateProductResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ProductUsecase.Create")
	defer span.End()

	err = uc.userRepository.CreateOrUpdate(ctx, entity.Product{
		Name:        req.Name,
		SKU:         req.SKU,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to save product"))
		return
	}

	resp.Name = req.Name
	resp.SKU = req.SKU

	return
}
