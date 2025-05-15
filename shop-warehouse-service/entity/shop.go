package entity

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Shop struct {
	UUID      uuid.UUID      `gorm:"column:uuid;->" json:"uuid"`
	Name      string         `gorm:"column:name" json:"name"`
	Address   null.String    `gorm:"column:address" json:"address"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

func (Shop) TableName() string {
	return "shops"
}
