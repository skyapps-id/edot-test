package driver

import (
	"fmt"
	"time"

	"github.com/skyapps-id/edot-test/shop-warehouse-service/config"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDatabase(cfg config.Config) *gorm.DB {
	dialector := postgres.Open(cfg.DbUrl)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to open gorm DB: %w", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get underlying sql.DB: %w", err))
	}

	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	if cfg.DbDebug {
		db = db.Debug()
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(fmt.Errorf("failed to use otelgorm plugin: %w", err))
	}

	return db
}
