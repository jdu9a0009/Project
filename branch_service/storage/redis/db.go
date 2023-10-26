package redis

import (
	"branch_service/config"
	"branch_service/storage"
	"context"

	"fmt"

	"github.com/go-redis/cache/v9"
	goRedis "github.com/redis/go-redis/v9"
)

type cacheStrg struct {
	db     *cache.Cache
	cacheR *cacheRepo
}

func NewCache(ctx context.Context, cfg config.Config) (storage.CacheI, error) {
	redisClient := goRedis.NewClient(&goRedis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDatabase,
	})

	redisCache := cache.New(&cache.Options{
		Redis: redisClient,
	})

	return &cacheStrg{
		db: redisCache,
	}, nil

}

func (r *cacheStrg) Cache() storage.RedisI {
	if r.cacheR == nil {
		r.cacheR = NewCacheRepo(r.db)
	}
	return r.cacheR
}
