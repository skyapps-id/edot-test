package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"

	"github.com/skyapps-id/edot-test/order-service/entity"
	"github.com/skyapps-id/edot-test/order-service/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Order interface {
	Create(ctx context.Context, order entity.Order, orderItems []entity.OrderItem) (tx *gorm.DB, err error)
	GetAll(ctx context.Context, name null.String, limit, offset int, sort string) (order []entity.Order, count int64, err error)
	FindByUUID(ctx context.Context, uuid uuid.UUID) (order entity.Order, err error)
	FindBySKU(ctx context.Context, sku string) (order entity.Order, err error)
	FindByName(ctx context.Context, name string) (order entity.Order, err error)
	UpdateStatus(ctx context.Context, uuid uuid.UUID, status string) (err error)
}

type order struct {
	tableName  string
	tableName1 string
	database   *gorm.DB
}

func NewOrderRepository(database *gorm.DB) Order {
	return &order{
		tableName:  entity.Order{}.TableName(),
		tableName1: entity.OrderItem{}.TableName(),
		database:   database,
	}
}

func (r *order) Create(ctx context.Context, order entity.Order, orderItems []entity.OrderItem) (tx *gorm.DB, err error) {
	tx = r.database.WithContext(ctx).Begin()
	err = tx.Table(r.tableName).Create(&order).Error
	if err != nil {
		logger.Log.Error("Error in OrderRepository.Create",
			zap.Error(err),
			zap.String("module", "OrderRepository"),
			zap.String("method", "CreateOrder"),
		)
		tx.Rollback()
	}

	err = tx.Table(r.tableName1).Create(&orderItems).Error
	if err != nil {
		logger.Log.Error("Error in OrderRepository.Create",
			zap.Error(err),
			zap.String("module", "OrderRepository"),
			zap.String("method", "CreateOrderItems"),
		)
		tx.Rollback()
	}

	return
}

func (r *order) GetAll(ctx context.Context, name null.String, limit, offset int, sort string) (orders []entity.Order, count int64, err error) {
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
		Find(&orders).Error
	if err != nil {
		logger.Log.Error("Error in OrderRepository.GetAll",
			zap.Error(err),
			zap.String("module", "OrderRepository"),
			zap.String("method", "GetAll"),
		)
	}

	return
}

func (r *order) FindByUUID(ctx context.Context, uuid uuid.UUID) (order entity.Order, err error) {
	err = r.database.WithContext(ctx).Where("uuid = ?", uuid).First(&order).Error
	if err != nil {
		logger.Log.Error("Error in OrderRepository.FindByUUID",
			zap.Error(err),
			zap.String("module", "OrderRepository"),
			zap.String("method", "FindByUUID"),
		)
	}

	return
}

func (r *order) FindBySKU(ctx context.Context, sku string) (order entity.Order, err error) {
	err = r.database.WithContext(ctx).Where("sku = ?", sku).First(&order).Error
	if err != nil {
		logger.Log.Error("Error in OrderRepository.FindBySKU",
			zap.Error(err),
			zap.String("module", "OrderRepository"),
			zap.String("method", "FindBySKU"),
		)
	}

	return
}

func (r *order) FindByName(ctx context.Context, name string) (order entity.Order, err error) {
	err = r.database.WithContext(ctx).Where("name LIKE ?", name).First(&order).Error
	if err != nil {
		logger.Log.Error("Error in OrderRepository.FindByName",
			zap.Error(err),
			zap.String("module", "OrderRepository"),
			zap.String("method", "FindByName"),
		)
	}

	return
}

func (r *order) UpdateStatus(ctx context.Context, uuid uuid.UUID, status string) (err error) {
	err = r.database.WithContext(ctx).Model(entity.Order{}).Where("uuid = ?", uuid).Update("status", status).Error
	if err != nil {
		logger.Log.Error("Error in OrderRepository.UpdateStatus",
			zap.Error(err),
			zap.String("module", "OrderRepository"),
			zap.String("method", "UpdateStatus"),
		)
	}

	return
}
