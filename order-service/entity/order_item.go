package entity

import "github.com/google/uuid"

type OrderItem struct {
	UUID          uuid.UUID `gorm:"column:uuid;->" json:"uuid"`
	OrderUUID     uuid.UUID `gorm:"column:order_uuid" json:"order_uuid"`
	ProductUUID   uuid.UUID `gorm:"column:product_uuid" json:"product_uuid"`
	StoreUUID     uuid.UUID `gorm:"column:store_uuid" json:"store_uuid"`
	WarehouseUUID uuid.UUID `gorm:"column:warehouse_uuid" json:"warehouse_uuid"`
	Quantity      int       `gorm:"column:quantity" json:"quantity"`
	ProductName   string    `gorm:"column:product_name" json:"product_name"`
	ProductSKU    string    `gorm:"column:product_sku" json:"product_sku"`
	Price         float64   `gorm:"column:price" json:"price"`
	TotalPrice    float64   `gorm:"column:total_price" json:"total_price"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
