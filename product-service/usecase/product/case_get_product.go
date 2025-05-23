package product

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/product-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/product-service/pkg/tracer"
	util "github.com/skyapps-id/edot-test/product-service/pkg/utils"
	"github.com/skyapps-id/edot-test/product-service/wrapper/shop_warehouse_service"
)

func (uc *usecase) Gets(ctx context.Context, req GetProductsRequest) (resp GetProductsResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ProductUsecase.Gets")
	defer span.End()

	products, count, err := uc.productRepository.GetAll(
		ctx,
		req.Name,
		int(req.PerPage.Int64), int(req.Page.Int64), req.Sort)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to gets product"))
		return
	}

	resp.List = uc.productsMapper(products)
	resp.Pagination = util.Pagination(int(req.Page.Int64), int(req.PerPage.Int64), count)

	return
}

func (uc *usecase) Get(ctx context.Context, req GetProductRequest) (resp GetProductResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ProductUsecase.Get")
	defer span.End()

	product, err := uc.productRepository.FindByUUID(ctx, req.UUID)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get product"))
		return
	}

	productStock, err := uc.shopWarehouseWrapper.GetProductStock(ctx, shop_warehouse_service.GetProductStockRequest{
		ProductUUID: req.UUID,
	})
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get product stock %w", err))
		return
	}
	resp = uc.productMapper(product)
	resp.Stock = productStock.Quantity

	return
}

func (uc *usecase) GetByUUIDs(ctx context.Context, req GetProductByUUIDsRequest) (resp map[uuid.UUID]GetProductResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ProductUsecase.GetByUUIDs")
	defer span.End()

	products, err := uc.productRepository.GetByUUIDs(ctx, req.UUIDs)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get product"))
		return
	}

	resp = uc.productByUUidsMapper(products)

	return
}
