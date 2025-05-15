package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/product-service/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product interface {
	CreateOrUpdate(ctx context.Context, user entity.Product) (err error)
	FindByUUID(ctx context.Context, uuid uuid.UUID) (user entity.Product, err error)
	FindBySKU(ctx context.Context, sku string) (user entity.Product, err error)
	FindByName(ctx context.Context, name string) (user entity.Product, err error)
}

type user struct {
	tableName string
	database  *gorm.DB
}

func NewUserRepository(database *gorm.DB) Product {
	return &user{
		tableName: entity.Product{}.TableName(),
		database:  database,
	}
}

func (r *user) CreateOrUpdate(ctx context.Context, user entity.Product) (err error) {
	err = r.database.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "sku"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name", "description", "price", "image_url",
		}),
	}).Create(&user).Error

	return
}

func (r *user) FindByUUID(ctx context.Context, uuid uuid.UUID) (user entity.Product, err error) {
	err = r.database.WithContext(ctx).Where("uuid = ?", uuid).First(&user).Error
	return
}

func (r *user) FindBySKU(ctx context.Context, sku string) (user entity.Product, err error) {
	err = r.database.WithContext(ctx).Where("sku = ?", sku).First(&user).Error
	return
}

func (r *user) FindByName(ctx context.Context, name string) (user entity.Product, err error) {
	err = r.database.WithContext(ctx).Where("name LIKE ?", name).First(&user).Error
	return
}
