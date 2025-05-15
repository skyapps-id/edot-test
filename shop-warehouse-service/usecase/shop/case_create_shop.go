package shop

import (
	"context"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/tracer"
)

func (uc *usecase) Craete(ctx context.Context, req CreateShopRequest) (resp CreateShopResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ShopUsecase.Create")
	defer span.End()

	err = uc.shopRepository.Create(ctx, entity.Shop{
		Name:    req.Name,
		Address: req.Address,
	})
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to save shop"))
		return
	}

	resp.Name = req.Name

	return
}
