package shop

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/tracer"
	util "github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/utils"
	"gorm.io/gorm"
)

func (uc *usecase) Gets(ctx context.Context, req GetShopsRequest) (resp GetShopsResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ShopUsecase.Gets")
	defer span.End()

	shops, count, err := uc.shopRepository.GetAll(
		ctx,
		req.Name,
		int(req.PerPage.Int64), int(req.Page.Int64), req.Sort)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to gets shop"))
		return
	}

	resp.List = uc.shopsMapper(shops)
	resp.Pagination = util.Pagination(int(req.Page.Int64), int(req.PerPage.Int64), count)

	return
}

func (uc *usecase) Get(ctx context.Context, req GetShopRequest) (resp GetShopResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ShopUsecase.Get")
	defer span.End()

	shop, err := uc.shopRepository.FindByUUID(ctx, req.UUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("shop not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get shop"))
		return
	}

	resp = uc.shopMapper(shop)

	return
}
