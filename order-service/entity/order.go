package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	UUID       uuid.UUID `gorm:"column:uuid" json:"uuid"`
	OrderID    string    `gorm:"column:order_id" json:"order_id"`
	UserUUID   uuid.UUID `gorm:"column:user_uuid" json:"user_uuid"`
	Status     string    `gorm:"column:status" json:"status"`
	TotalItems int       `gorm:"column:total_items" json:"total_items"`
	TotalPrice float64   `gorm:"column:total_price" json:"total_price"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}
