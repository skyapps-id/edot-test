package order

import (
	"context"
	"fmt"
	"net/http"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/skyapps-id/edot-test/order-service/entity"
	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
)

func (uc *usecase) Craete(ctx context.Context, req CreateOrderRequest) (resp CreateOrderResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "OrderUsecase.Create")
	defer span.End()

	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	id, _ := gonanoid.Generate(alphabet, 16)
	order := entity.Order{
		UserUUID:   req.UserUUID,
		Status:     "checkout",
		OrderID:    "TX-" + id,
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
