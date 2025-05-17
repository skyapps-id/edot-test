package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/order-service/entity"
	"github.com/skyapps-id/edot-test/order-service/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderItem interface {
	GetByOrderUUID(ctx context.Context, orderUUID uuid.UUID) (orderItems []entity.OrderItem, err error)
}

type orderItem struct {
	tableName string
	database  *gorm.DB
}

func NewOrderItemRepository(database *gorm.DB) OrderItem {
	return &orderItem{
		tableName: entity.OrderItem{}.TableName(),
		database:  database,
	}
}

func (r *orderItem) GetByOrderUUID(ctx context.Context, orderUUID uuid.UUID) (orderItems []entity.OrderItem, err error) {
	err = r.database.WithContext(ctx).Where("order_uuid = ?", orderUUID).Find(&orderItems).Error
	if err != nil {
		logger.Log.Error("Error in OrderItemRepository.GetByOrderUUID",
			zap.Error(err),
			zap.String("module", "OrderItemRepository"),
			zap.String("method", "GetByOrderUUID"),
		)
	}

	return
}
