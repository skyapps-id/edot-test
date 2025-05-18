package task

import (
	"context"

	"github.com/skyapps-id/edot-test/order-service/pkg/logger"
	"github.com/skyapps-id/edot-test/order-service/usecase/order"
	"go.uber.org/zap"
)

type PaymentTask interface {
	OrderCheck(b64payload string) (err error)
}

type task struct {
	orderUsecase order.OrderUsecase
}

func NewWrapper(orderUsecase order.OrderUsecase) PaymentTask {
	return &task{
		orderUsecase: orderUsecase,
	}
}

func (t *task) OrderCheck(b64payload string) (err error) {
	var payload order.OrderCancelRequest
	DecodeToTask(b64payload, &payload)

	_, err = t.orderUsecase.OrderCancel(context.Background(), payload)
	if err != nil {
		logger.Log.Info("fail order cancel task", zap.Error(err))
		return
	}

	return
}
