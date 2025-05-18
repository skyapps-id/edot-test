package order

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
	"github.com/skyapps-id/edot-test/order-service/wrapper/shop_warehouse_service"
	"gorm.io/gorm"
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

func (uc *usecase) OrderCancel(ctx context.Context, req OrderCancelRequest) (resp OrderCancelResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "OrderUsecase.OrderCancel")
	defer span.End()

	order, err := uc.orderRepository.FindByUUID(ctx, req.OrderUUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("shop not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get order"))
		return
	}

	if order.Status == "payment" {
		return
	}

	err = uc.orderRepository.UpdateStatus(ctx, req.OrderUUID, "cancel")
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to update order to cancel"))
		return
	}

	var payload shop_warehouse_service.ProductStockAdditionRequest
	for _, row := range req.Products {
		payload.Products = append(payload.Products, shop_warehouse_service.DataProductStock{
			ProductUUID:   row.ProductUUID,
			WarehouseUUID: row.WarehouseUUID,
			Quantity:      row.Quantity,
		})
	}
	_, err = uc.shopWarehouseWrapper.ProductStockAddition(ctx, payload)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail product stock reduction: %w", err))
		return
	}

	return
}
