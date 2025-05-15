package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/order-service/entity"
	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
	util "github.com/skyapps-id/edot-test/order-service/pkg/utils"
)

func (uc *usecase) Gets(ctx context.Context, req GetOrdersRequest) (resp GetOrdersResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "OrderUsecase.Gets")
	defer span.End()

	orders, count, err := uc.orderRepository.GetAll(
		ctx,
		req.Name,
		int(req.PerPage.Int64), int(req.Page.Int64), req.Sort)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to gets order"))
		return
	}

	resp.List = uc.ordersMapper(orders)
	resp.Pagination = util.Pagination(int(req.Page.Int64), int(req.PerPage.Int64), count)

	return
}

func (uc *usecase) Get(ctx context.Context, req GetOrderRequest) (resp GetOrderResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "OrderUsecase.Get")
	defer span.End()

	order, err := uc.orderRepository.FindByUUID(ctx, req.UUID)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get order"))
		return
	}

	resp = uc.orderMapper(order, []entity.OrderItem{})

	return
}
