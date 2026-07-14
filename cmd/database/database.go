package database

import (
	"client-server/cmd/models"
	"context"
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "file:database.sqlite?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	db.AutoMigrate(models.CurrencyInfo{})
	return db, nil
}

func Insert(currency *models.CurrencyInfo) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	return db.WithContext(ctx).Create(currency).Error
}
