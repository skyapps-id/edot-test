package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	UUID        uuid.UUID      `gorm:"column:uuid;->" json:"uuid"`
	Name        string         `gorm:"column:name" json:"name"`
	SKU         string         `gorm:"column:sku" json:"sku"`
	Description string         `gorm:"column:description" json:"description"`
	Price       float64        `gorm:"column:price" json:"price"`
	ImageURL    string         `gorm:"column:image_url" json:"image_url"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (Product) TableName() string {
	return "products"
}
