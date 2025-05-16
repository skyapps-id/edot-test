package repository

import (
	"context"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type WarehouseProduct interface {
	Create(ctx context.Context, warehouseProduct entity.WarehouseProduct) (err error)
}

type warehouseProduct struct {
	tableName string
	database  *gorm.DB
}

func NewWarehouseProductRepository(database *gorm.DB) WarehouseProduct {
	return &warehouseProduct{
		tableName: entity.WarehouseProduct{}.TableName(),
		database:  database,
	}
}

func (r *warehouseProduct) Create(ctx context.Context, warehouseProduct entity.WarehouseProduct) (err error) {
	err = r.database.WithContext(ctx).Create(&warehouseProduct).Error
	if err != nil {
		logger.Log.Error("Error in WarehouseProductRepository.Create",
			zap.Error(err),
			zap.String("module", "WarehouseProductRepository"),
			zap.String("method", "Create"),
		)
	}
	return
}
