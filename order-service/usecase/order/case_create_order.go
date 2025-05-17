package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/skyapps-id/edot-test/order-service/entity"
	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/constant"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
	"github.com/skyapps-id/edot-test/order-service/wrapper/product_service"
	"github.com/skyapps-id/edot-test/order-service/wrapper/shop_warehouse_service"
)

func (uc *usecase) Craete(ctx context.Context, req CreateOrderRequest) (resp CreateOrderResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "OrderUsecase.Create")
	defer span.End()

	products, err := uc.productWrapper.GetProducts(ctx, product_service.ProductRequest{
		Uuids: req.GetProductUUIDs(),
	})
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get product %w", err))
		return
	}

	productStock, err := uc.shopWarehouseWrapper.GetProductStock(ctx, shop_warehouse_service.ProductStockRequest{
		Uuids: req.GetProductUUIDs(),
	})
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get shop warehouse: %w", err))
		return
	}

	id, _ := gonanoid.Generate(constant.ALPHABETNUMBER, 16)
	order := entity.Order{
		UUID:     uuid.New(),
		UserUUID: req.UserUUID,
		Status:   "checkout",
		OrderID:  "TX-" + id,
	}

	orderItems := []entity.OrderItem{}
	totalPrice := float64(0)
	productReduction := shop_warehouse_service.ProductStockReductionRequest{}
	totaItem := 0
	for _, row := range req.Orderitems {
		product, ok := products[row.ProductUUID]
		if !ok {
			err = apperror.New(
				http.StatusUnprocessableEntity,
				fmt.Errorf("product not found uuid: %s", row.ProductUUID),
			)
			return
		}

		shopWarehouse, ok := productStock[row.ProductUUID]
		if !ok {
			err = apperror.New(
				http.StatusUnprocessableEntity,
				fmt.Errorf("product stock not found: [%s,%s,%s]", row.ProductUUID, product.SKU, product.Name),
			)
			return
		}

		if shopWarehouse.Quantity < row.Quantity {
			err = apperror.New(
				http.StatusUnprocessableEntity,
				fmt.Errorf("insufficient stock for product: [%s,%s,%s]", row.ProductUUID, product.SKU, product.Name),
			)
			return
		}

		totalPriceItem := product.Price * float64(row.Quantity)
		orderItems = append(orderItems, entity.OrderItem{
			OrderUUID:     order.UUID,
			StoreUUID:     shopWarehouse.ShopUUID,
			WarehouseUUID: shopWarehouse.WarehouseUUID,
			ProductUUID:   product.UUID,
			ProductName:   product.Name,
			ProductSKU:    product.SKU,
			Quantity:      row.Quantity,
			Price:         product.Price,
			TotalPrice:    totalPriceItem,
		})
		totalPrice += totalPriceItem
		totaItem += row.Quantity

		productReduction.Products = append(productReduction.Products, shop_warehouse_service.DataProductStock{
			ProductUUID:   product.UUID,
			WarehouseUUID: shopWarehouse.WarehouseUUID,
			Quantity:      row.Quantity,
		})
	}
	order.TotalItems = totaItem

	tx, err := uc.orderRepository.Create(ctx, order, orderItems)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to save order"))
		return
	}

	_, err = uc.shopWarehouseWrapper.ProductStockReduction(ctx, productReduction)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail product stock reduction: %w", err))
		tx.Rollback()
		return
	}

	tx.Commit()
	resp.OrderID = order.OrderID

	return
}
