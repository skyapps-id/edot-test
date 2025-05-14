package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/user-service/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User interface {
	CreateOrUpdate(ctx context.Context, user entity.User) (err error)
	FindByUUID(ctx context.Context, uuid uuid.UUID) (user entity.User, err error)
	FindByEmailOrPhone(ctx context.Context, email, phone string) (entity.User, error)
}

type user struct {
	tableName string
	database  *gorm.DB
}

func NewUserRepository(database *gorm.DB) User {
	return &user{
		tableName: "users",
		database:  database,
	}
}

func (r *user) CreateOrUpdate(ctx context.Context, user entity.User) (err error) {
	err = r.database.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"phone", "password", "name"}),
	}).Create(&user).Error

	return
}

func (r *user) FindByUUID(ctx context.Context, uuid uuid.UUID) (user entity.User, err error) {
	err = r.database.WithContext(ctx).Where("uuid = ?", uuid).First(&user).Error
	return
}

func (r *user) FindByEmailOrPhone(ctx context.Context, email, phone string) (user entity.User, err error) {
	err = r.database.WithContext(ctx).Where("email = ? OR phone = ?", email, phone).First(&user).Error
	return
}
