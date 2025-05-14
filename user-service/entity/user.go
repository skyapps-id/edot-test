package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID `gorm:"column:uuid;->" json:"uuid"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email" json:"email"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
