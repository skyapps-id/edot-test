package order

import (
	"context"
	"fmt"
	"net/http"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/skyapps-id/edot-test/order-service/entity"
	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
	"github.com/skyapps-id/edot-test/order-service/wrapper/product_service"
)

func (uc *usecase) Craete(ctx context.Context, req CreateOrderRequest) (resp CreateOrderResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "OrderUsecase.Create")
	defer span.End()
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	products, err := uc.productWrapper.GetProducts(ctx, product_service.ProductRequest{
		Uuids: req.GetProductUUIDs(),
	})

	id, _ := gonanoid.Generate(alphabet, 16)
	order := entity.Order{
		UserUUID:   req.UserUUID,
		Status:     "checkout",
		OrderID:    "TX-" + id,
		TotalItems: 1,
		TotalPrice: 1,
	}

	orderItems := []entity.OrderItem{}
	for _, row := range products {
		orderItems = append(orderItems, entity.OrderItem{
			ProductUUID: row.UUID,
			ProductName: row.Name,
			ProductSKU:  row.SKU,
			Quantity:    0,
			Price:       row.Price,
		})
	}

	fmt.Println(orderItems)

	err = uc.orderRepository.Create(ctx, order)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to save order"))
		return
	}

	resp.OrderID = order.OrderID

	return
}
