package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/order-service/entity"
	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
)

func (uc *usecase) Craete(ctx context.Context, req CreateOrderRequest) (resp CreateOrderResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "OrderUsecase.Create")
	defer span.End()

	order := entity.Order{
		UserUUID:   uuid.New(),
		Status:     "checkout",
		TotalItems: 1,
		TotalPrice: 1,
	}

	err = uc.orderRepository.Create(ctx, order)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to save order"))
		return
	}

	resp.OrderID = order.OrderID

	return
}
