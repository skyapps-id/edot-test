package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type WarehouseProduct interface {
	Create(ctx context.Context, warehouseProduct entity.WarehouseProduct) (err error)
	GetMaxQuantityByProductUUIDs(ctx context.Context, productUUIDs []uuid.UUID) (warehouse []entity.WarehouseProduct, err error)
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

func (r *warehouseProduct) GetMaxQuantityByProductUUIDs(ctx context.Context, productUUIDs []uuid.UUID) (warehouse []entity.WarehouseProduct, err error) {
	err = r.database.WithContext(ctx).
		Select(`
			DISTINCT ON (warehouse_products.product_uuid)
			warehouse_products.uuid,
			warehouse_products.warehouse_uuid,
			warehouse_products.product_uuid,
			warehouse_products.quantity,
			warehouse_products.updated_at,
			warehouses.shop_uuid AS shop_uuid
		`).
		Joins("JOIN warehouses ON warehouses.uuid = warehouse_products.warehouse_uuid AND warehouses.active = true").
		Where("warehouse_products.product_uuid IN ?", productUUIDs).
		Order("warehouse_products.product_uuid, warehouse_products.quantity DESC").
		Find(&warehouse).Error
	if err != nil {
		logger.Log.Error("Error in WarehouseProductRepository.GetMaxStockByProductUUID",
			zap.Error(err),
			zap.String("module", "WarehouseProductRepository"),
			zap.String("method", "GetMaxStockByProductUUID"),
		)
	}

	return
}
