package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warehouse struct {
	UUID      uuid.UUID      `gorm:"column:uuid;->" json:"uuid"`
	Name      string         `gorm:"column:name" json:"name"`
	Address   string         `gorm:"column:address" json:"address"`
	Active    bool           `gorm:"column:active" json:"active"`
	ShopUUID  uuid.UUID      `gorm:"column:shop_uuid" json:"shop_uuid"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

func (Warehouse) TableName() string {
	return "warehouses"
}
