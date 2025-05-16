package entity

import (
	"time"

	"github.com/google/uuid"
)

type WarehouseProduct struct {
	UUID          uuid.UUID `gorm:"column:uuid;->" json:"uuid"`
	WarehouseUUID uuid.UUID `gorm:"column:warehouse_uuid" json:"warehouse_uuid"`
	ProductUUID   uuid.UUID `gorm:"column:product_uuid" json:"product_uuid"`
	Quantity      int       `gorm:"column:quantity" json:"quantity"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	ShopUUID      uuid.UUID `gorm:"column:shop_uuid;->" json:"shop_uuid"`
}

func (WarehouseProduct) TableName() string {
	return "warehouse_products"
}
