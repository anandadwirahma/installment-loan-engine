package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var AppConfig Config

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("Error loading .env file: %v", err))
	}

	err = env.Parse(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("Error parsing env: %v", err))
	}
}

func GetEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultVal
}
