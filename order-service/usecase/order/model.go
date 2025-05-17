package order

import (
	"time"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/order-service/pkg/response"
	"gopkg.in/guregu/null.v4"
)

type OrderItemsRequest struct {
	ProductUUID uuid.UUID `json:"product_uuid" validate:"required"`
	Quantity    int       `json:"quantity" validate:"required"`
}

type CreateOrderRequest struct {
	UserUUID   uuid.UUID           `json:"user_uuid"`
	Orderitems []OrderItemsRequest `json:"order_items" validate:"required"`
}

func (t CreateOrderRequest) GetProductUUIDs() (uuids []uuid.UUID) {
	for _, row := range t.Orderitems {
		uuids = append(uuids, row.ProductUUID)
	}

	return
}

type CreateOrderResponse struct {
	OrderID string `json:"order_id"`
}

type GetOrdersRequest struct {
	Name    null.String `query:"name"`
	Page    null.Int    `query:"page"`
	PerPage null.Int    `query:"per_page"`
	Sort    string      `query:"sort" validate:"oneof='ASC' 'DESC'"`
}

type DataOrders struct {
	UUID       uuid.UUID `json:"uuid"`
	OrderID    string    `json:"order_id"`
	Status     string    `json:"status"`
	TotalItems int       `json:"total_items"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetOrdersResponse struct {
	List       []DataOrders        `json:"list"`
	Pagination response.Pagination `json:"pagination"`
}

type GetOrderRequest struct {
	UUID uuid.UUID `param:"uuid"`
}

type OrderItemsResponse struct {
	ProductUUID uuid.UUID `json:"product_uuid"`
	Quantity    int       `json:"quantity"`
	ProductName string    `json:"product_name"`
	ProductSKU  string    `json:"product_sku"`
	Price       float64   `json:"price"`
}

type GetOrderResponse struct {
	UUID       uuid.UUID            `json:"uuid"`
	OrderID    string               `json:"order_id"`
	Status     string               `json:"status"`
	TotalItems int                  `json:"total_items"`
	TotalPrice float64              `json:"total_price"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
	OrderItems []OrderItemsResponse `json:"order_items"`
}

type OrderStatusToPaymentRequest struct {
	UUID uuid.UUID `param:"uuid"`
}

type OrderStatusToPaymentResponse struct {
	UUID uuid.UUID `json:"uuid"`
}
