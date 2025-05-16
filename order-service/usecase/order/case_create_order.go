package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/skyapps-id/edot-test/order-service/entity"
	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
	"github.com/skyapps-id/edot-test/order-service/wrapper/product_service"
)

func (uc *usecase) Craete(ctx context.Context, req CreateOrderRequest) (resp CreateOrderResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "OrderUsecase.Create")
	defer span.End()

	products, err := uc.productWrapper.GetProducts(ctx, product_service.ProductRequest{
		Uuids: req.GetProductUUIDs(),
	})

	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id, _ := gonanoid.Generate(alphabet, 16)
	order := entity.Order{
		UUID:     uuid.New(),
		UserUUID: req.UserUUID,
		Status:   "checkout",
		OrderID:  "TX-" + id,
	}

	orderItems := []entity.OrderItem{}
	totalPrice := float64(0)
	totaItem := 0
	for _, row := range req.Orderitems {
		product := products[row.ProductUUID]
		totalPriceItem := product.Price * float64(row.Quantity)
		orderItems = append(orderItems, entity.OrderItem{
			OrderUUID:   order.UUID,
			ProductUUID: product.UUID,
			ProductName: product.Name,
			ProductSKU:  product.SKU,
			Quantity:    row.Quantity,
			Price:       product.Price,
			TotalPrice:  totalPriceItem,
		})
		totalPrice += totalPriceItem
		totaItem += row.Quantity
	}
	order.TotalItems = totaItem

	err = uc.orderRepository.Create(ctx, order, orderItems)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to save order"))
		return
	}

	resp.OrderID = order.OrderID

	return
}
