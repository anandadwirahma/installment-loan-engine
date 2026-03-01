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
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("Error loading .env file: %v", err))
	}

	err = env.Parse(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("Error parsing env: %v", err))
	}

	logger.Infof("Config loaded: %+v\n", AppConfig)
}

func GetEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultVal
}
