package task

import "github.com/google/uuid"

type DataProductStock struct {
	ProductUUID   uuid.UUID `json:"product_uuid"`
	WarehouseUUID uuid.UUID `json:"warehouse_uuid"`
	Quantity      int       `json:"quantity"`
}

type TaskOrderCancel struct {
	OrderUUID uuid.UUID          `json:"order_uuid"`
	Products  []DataProductStock `json:"products"`
}
