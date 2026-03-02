package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"installment-loan-engine/internal/shared/config"
	"installment-loan-engine/internal/shared/logger"
)

var RDB *redis.Client

func Init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.AppConfig.Redis.Host, config.AppConfig.Redis.Port),
		Password: config.AppConfig.Redis.Password,
		DB:       config.AppConfig.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RDB.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed to connect Redis: %v", err))
	}

	logger.Infof("Redis connection established")
}
