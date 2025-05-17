package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Warehouse interface {
	Create(ctx context.Context, warehouse entity.Warehouse) (err error)
	GetAll(ctx context.Context, name null.String, limit, offset int, sort string) (shop []entity.Warehouse, count int64, err error)
	FindByUUID(ctx context.Context, uuid uuid.UUID) (warehouse entity.Warehouse, err error)
	WarehouseUpdateActive(ctx context.Context, uuid uuid.UUID, isActive bool) (err error)
}

type warehouse struct {
	tableName string
	database  *gorm.DB
}

func NewWarehouseRepository(database *gorm.DB) Warehouse {
	return &warehouse{
		tableName: entity.Warehouse{}.TableName(),
		database:  database,
	}
}

func (r *warehouse) Create(ctx context.Context, warehouse entity.Warehouse) (err error) {
	err = r.database.WithContext(ctx).Create(&warehouse).Error
	if err != nil {
		logger.Log.Error("Error in WarehouseRepository.Create",
			zap.Error(err),
			zap.String("module", "WarehouseRepository"),
			zap.String("method", "Create"),
		)
	}
	return
}

func (r *warehouse) GetAll(ctx context.Context, name null.String, limit, offset int, sort string) (products []entity.Warehouse, count int64, err error) {
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
		logger.Log.Error("Error in WarehouseRepository.GetAll",
			zap.Error(err),
			zap.String("module", "WarehouseRepository"),
			zap.String("method", "GetAll"),
		)
	}

	return
}

func (r *warehouse) FindByUUID(ctx context.Context, uuid uuid.UUID) (warehouse entity.Warehouse, err error) {
	err = r.database.WithContext(ctx).Where("uuid = ?", uuid).First(&warehouse).Error
	if err != nil {
		logger.Log.Error("Error in WarehouseRepository.FindByUUID",
			zap.Error(err),
			zap.String("module", "WarehouseRepository"),
			zap.String("method", "FindByUUID"),
		)
	}

	return
}

func (r *warehouse) WarehouseUpdateActive(ctx context.Context, uuid uuid.UUID, isActive bool) (err error) {
	err = r.database.WithContext(ctx).Model(entity.Warehouse{}).Where("uuid = ?", uuid).Update("active", isActive).Error
	if err != nil {
		logger.Log.Error("Error in WarehouseRepository.WarehouseUpdateActive",
			zap.Error(err),
			zap.String("module", "WarehouseRepository"),
			zap.String("method", "WarehouseUpdateActive"),
		)
	}

	return
}
