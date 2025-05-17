package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WarehouseProduct interface {
	Create(ctx context.Context, warehouseProduct entity.WarehouseProduct) (err error)
	GetMaxQuantityByProductUUIDs(ctx context.Context, productUUIDs []uuid.UUID) (warehouse []entity.WarehouseProduct, err error)
	GetProductStock(ctx context.Context, productUUID uuid.UUID) (warehouseProduct entity.WarehouseProduct, err error)
	ProductStockAddition(ctx context.Context, products []entity.ProductStock) (err error)
	ProductStockReduction(ctx context.Context, products []entity.ProductStock) (err error)
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

func (r *warehouseProduct) GetMaxQuantityByProductUUIDs(ctx context.Context, productUUIDs []uuid.UUID) (warehouseProducts []entity.WarehouseProduct, err error) {
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
		Find(&warehouseProducts).Error
	if err != nil {
		logger.Log.Error("Error in WarehouseProductRepository.GetMaxStockByProductUUID",
			zap.Error(err),
			zap.String("module", "WarehouseProductRepository"),
			zap.String("method", "GetMaxStockByProductUUID"),
		)
	}

	return
}

func (r *warehouseProduct) GetProductStock(ctx context.Context, productUUID uuid.UUID) (warehouseProduct entity.WarehouseProduct, err error) {
	err = r.database.WithContext(ctx).
		Select(`
			warehouse_products.uuid,
			warehouse_products.warehouse_uuid,
			warehouse_products.product_uuid,
			SUM(warehouse_products.quantity) AS quantity,
			warehouse_products.updated_at,
		`).
		Joins("JOIN warehouses ON warehouses.uuid = warehouse_products.warehouse_uuid AND warehouses.active = true").
		Where("warehouse_products.product_uuid = ?", productUUID).
		First(&warehouseProduct).Error
	if err != nil {
		logger.Log.Error("Error in WarehouseProductRepository.GetAllStock",
			zap.Error(err),
			zap.String("module", "WarehouseProductRepository"),
			zap.String("method", "GetAllStock"),
		)
	}

	return
}

func (r *warehouseProduct) ProductStockAddition(ctx context.Context, products []entity.ProductStock) (err error) {
	tx := r.database.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, row := range products {
		var warehouse entity.Warehouse
		err = r.database.WithContext(ctx).Where("uuid = ? AND active = true", row.WarehouseUUID).First(&warehouse).Error
		if err != nil {
			tx.Rollback()
			return apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("warehouse inactive"))
		}

		var warehouseProduct entity.WarehouseProduct
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("product_uuid = ? AND warehouse_uuid = ?", row.ProductUUID, row.WarehouseUUID).
			First(&warehouseProduct).Error
		if err != nil {
			tx.Rollback()
			return apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("product stock not found: %w", err))
		}

		err = tx.Model(&warehouseProduct).
			Where("product_uuid = ? AND warehouse_uuid = ?", row.ProductUUID, row.WarehouseUUID).
			Update("quantity", warehouseProduct.Quantity+row.Quantity).Error
		if err != nil {
			tx.Rollback()
			return apperror.New(http.StatusInternalServerError, fmt.Errorf("failed to update stock: %w", err))
		}
	}

	return tx.Commit().Error
}

func (r *warehouseProduct) ProductStockReduction(ctx context.Context, products []entity.ProductStock) (err error) {
	tx := r.database.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, row := range products {
		var warehouse entity.Warehouse
		err = r.database.WithContext(ctx).Where("uuid = ? AND active = true", row.WarehouseUUID).First(&warehouse).Error
		if err != nil {
			tx.Rollback()
			return apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("warehouse inactive"))
		}

		var warehouseProduct entity.WarehouseProduct
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("product_uuid = ? AND warehouse_uuid = ?", row.ProductUUID, row.WarehouseUUID).
			First(&warehouseProduct).Error
		if err != nil {
			tx.Rollback()
			return apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("product stock not found: %w", err))
		}

		if warehouseProduct.Quantity < row.Quantity {
			tx.Rollback()
			return apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("insufficient stock for product: %s", row.ProductUUID))
		}

		err = tx.Model(&warehouseProduct).
			Where("product_uuid = ? AND warehouse_uuid = ?", row.ProductUUID, row.WarehouseUUID).
			Update("quantity", warehouseProduct.Quantity-row.Quantity).Error
		if err != nil {
			tx.Rollback()
			return apperror.New(http.StatusInternalServerError, fmt.Errorf("failed to update stock: %w", err))
		}
	}

	return tx.Commit().Error
}
