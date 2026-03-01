package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"installment-loan-engine/internal/entity"
	"installment-loan-engine/internal/shared/config"
	"installment-loan-engine/internal/shared/logger"
)

var DB *gorm.DB

func Init() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Name,
		config.AppConfig.Database.Port,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get sql.DB: %v", err))
	}

	maxIdle := config.AppConfig.Database.MaxIdleConns
	maxOpen := config.AppConfig.Database.MaxOpenConns
	maxLifetime := config.AppConfig.Database.ConnMaxLifetimeMinutes

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)

	if err := DB.AutoMigrate(
		&entity.Loan{},
		&entity.Installment{},
		&entity.Transaction{},
	); err != nil {
		panic(fmt.Sprintf("AutoMigrate failed: %v", err))
	}

	logger.Infof("Database connection established (Idle: %d, Open: %d, Lifetime: %dm)\n", maxIdle, maxOpen, maxLifetime)
}
