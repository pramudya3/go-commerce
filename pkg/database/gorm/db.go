package database

import (
	"fmt"
	"go-commerce/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPass,
			cfg.DBName,
		),
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// Set up connection pool
	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)

	return database, nil
}
