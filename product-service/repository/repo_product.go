package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"

	"github.com/skyapps-id/edot-test/product-service/entity"
	"github.com/skyapps-id/edot-test/product-service/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product interface {
	CreateOrUpdate(ctx context.Context, product entity.Product) (err error)
	GetAll(ctx context.Context, name null.String, limit, offset int, sort string) (product []entity.Product, count int64, err error)
	FindByUUID(ctx context.Context, uuid uuid.UUID) (product entity.Product, err error)
	FindBySKU(ctx context.Context, sku string) (product entity.Product, err error)
	FindByName(ctx context.Context, name string) (product entity.Product, err error)
}

type product struct {
	tableName string
	database  *gorm.DB
}

func NewProductRepository(database *gorm.DB) Product {
	return &product{
		tableName: entity.Product{}.TableName(),
		database:  database,
	}
}

func (r *product) CreateOrUpdate(ctx context.Context, product entity.Product) (err error) {
	err = r.database.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "sku"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name", "description", "price", "image_url",
		}),
	}).Create(&product).Error
	if err != nil {
		logger.Log.Error("Error in ProductRepository.CreateOrUpdate",
			zap.Error(err),
			zap.String("module", "ProductRepository"),
			zap.String("method", "CreateOrUpdate"),
		)
	}

	return
}

func (r *product) GetAll(ctx context.Context, name null.String, limit, offset int, sort string) (products []entity.Product, count int64, err error) {
	query := r.database.WithContext(ctx).Table(r.tableName)

	if name.Valid {
		query = query.Where("name ILIKE ?", "%"+name.String+"%")
	}

	query.Select("uuid").Count(&count)

	err = query.
		Select("*").
		Limit(limit).
		Offset((offset - 1) * limit).
		Order(fmt.Sprintf("%s %s", "created_at", sort)).
		Find(&products).Error
	if err != nil {
		logger.Log.Error("Error in ProductRepository.GetAll",
			zap.Error(err),
			zap.String("module", "ProductRepository"),
			zap.String("method", "GetAll"),
		)
	}

	return
}

func (r *product) FindByUUID(ctx context.Context, uuid uuid.UUID) (product entity.Product, err error) {
	err = r.database.WithContext(ctx).Where("uuid = ?", uuid).First(&product).Error
	if err != nil {
		logger.Log.Error("Error in ProductRepository.FindByUUID",
			zap.Error(err),
			zap.String("module", "ProductRepository"),
			zap.String("method", "FindByUUID"),
		)
	}

	return
}

func (r *product) FindBySKU(ctx context.Context, sku string) (product entity.Product, err error) {
	err = r.database.WithContext(ctx).Where("sku = ?", sku).First(&product).Error
	if err != nil {
		logger.Log.Error("Error in ProductRepository.FindBySKU",
			zap.Error(err),
			zap.String("module", "ProductRepository"),
			zap.String("method", "FindBySKU"),
		)
	}

	return
}

func (r *product) FindByName(ctx context.Context, name string) (product entity.Product, err error) {
	err = r.database.WithContext(ctx).Where("name LIKE ?", name).First(&product).Error
	if err != nil {
		logger.Log.Error("Error in ProductRepository.FindByName",
			zap.Error(err),
			zap.String("module", "ProductRepository"),
			zap.String("method", "FindByName"),
		)
	}

	return
}
