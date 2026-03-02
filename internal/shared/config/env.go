package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	"installment-loan-engine/internal/shared/logger"
)

var AppConfig Config

func Init() {
	if err := godotenv.Load(); err != nil {
		logger.Infof("No .env file found, using system environment")
	}

	if err := env.Parse(&AppConfig); err != nil {
		panic(fmt.Errorf("error parsing env: %v", err))
	}

	logger.Infof("Config loaded: %+v\n", AppConfig)
}

func GetEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultVal
}
