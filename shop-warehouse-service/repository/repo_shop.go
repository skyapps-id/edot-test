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

type Shop interface {
	Create(ctx context.Context, shop entity.Shop) (err error)
	GetAll(ctx context.Context, name null.String, limit, offset int, sort string) (shop []entity.Shop, count int64, err error)
	FindByUUID(ctx context.Context, uuid uuid.UUID) (shop entity.Shop, err error)
}

type shop struct {
	tableName string
	database  *gorm.DB
}

func NewShopRepository(database *gorm.DB) Shop {
	return &shop{
		tableName: entity.Shop{}.TableName(),
		database:  database,
	}
}

func (r *shop) Create(ctx context.Context, shop entity.Shop) (err error) {
	err = r.database.WithContext(ctx).Create(&shop).Error
	if err != nil {
		logger.Log.Error("Error in ShopRepository.FindByUUID",
			zap.Error(err),
			zap.String("module", "ShopRepository"),
			zap.String("method", "Create"),
		)
	}
	return
}

func (r *shop) GetAll(ctx context.Context, name null.String, limit, offset int, sort string) (products []entity.Shop, count int64, err error) {
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
		logger.Log.Error("Error in ShopRepository.GetAll",
			zap.Error(err),
			zap.String("module", "ShopRepository"),
			zap.String("method", "GetAll"),
		)
	}

	return
}

func (r *shop) FindByUUID(ctx context.Context, uuid uuid.UUID) (shop entity.Shop, err error) {
	err = r.database.WithContext(ctx).Where("uuid = ?", uuid).First(&shop).Error
	if err != nil {
		logger.Log.Error("Error in ShopRepository.FindByUUID",
			zap.Error(err),
			zap.String("module", "ShopRepository"),
			zap.String("method", "FindByUUID"),
		)
	}

	return
}
