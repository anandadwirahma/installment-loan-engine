package repositories

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type CacheRepository interface {
	IncrPay(ctx context.Context, loanRefNum string) (int64, error)
	DelPayIncr(ctx context.Context, loanRefNum string) error
}

type cacheRepository struct {
	rdb *redis.Client
}

func NewCacheRepository(rdb *redis.Client) CacheRepository {
	return &cacheRepository{rdb: rdb}
}

func (c *cacheRepository) IncrPay(ctx context.Context, loanRefNum string) (int64, error) {
	key := fmt.Sprintf("loan#payment#%s", loanRefNum)
	return c.rdb.Incr(ctx, key).Result()
}

func (c *cacheRepository) DelPayIncr(ctx context.Context, loanRefNum string) error {
	key := fmt.Sprintf("loan#payment#%s", loanRefNum)
	return c.rdb.Del(ctx, key).Err()
}
