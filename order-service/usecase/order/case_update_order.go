package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
)

func (uc *usecase) UpdateStatusToPayment(ctx context.Context, req OrderStatusToPaymentRequest) (resp OrderStatusToPaymentResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "OrderUsecase.UpdateStatusToPayment")
	defer span.End()

	err = uc.orderRepository.UpdateStatus(ctx, req.UUID, "payment")
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to update order to payment"))
		return
	}

	resp.UUID = req.UUID

	return
}
