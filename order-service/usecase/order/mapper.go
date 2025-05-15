package order

import (
	"github.com/skyapps-id/edot-test/order-service/entity"
)

func (uc *usecase) ordersMapper(results []entity.Order) (resp []DataOrders) {
	for _, row := range results {
		order := DataOrders{
			UUID:       row.UUID,
			OrderID:    row.OrderID,
			Status:     row.Status,
			TotalItems: row.TotalItems,
			TotalPrice: row.TotalPrice,
			CreatedAt:  row.CreatedAt,
			UpdatedAt:  row.UpdatedAt,
		}
		resp = append(resp, order)
	}

	return
}

func (uc *usecase) orderMapper(order entity.Order, items []entity.OrderItem) (resp GetOrderResponse) {
	resp = GetOrderResponse{
		UUID:       order.UUID,
		OrderID:    order.OrderID,
		Status:     order.Status,
		TotalItems: order.TotalItems,
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}

	for _, row := range items {
		items := OrderItemsResponse{
			ProductUUID: row.ProductUUID,
			ProductSKU:  row.ProductSKU,
			ProductName: row.ProductName,
			Quantity:    row.Quantity,
			Price:       row.Price,
		}
		resp.OrderItems = append(resp.OrderItems, items)
	}

	return
}
