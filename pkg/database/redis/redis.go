package redis

import (
	"context"
	"go-commerce/pkg/config"
	"log"
	"time"

	rds "github.com/redis/go-redis/v9"
)

func NewRedis(config *config.Config) (*rds.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.TimeoutCtx)*time.Second)
	defer cancel()

	rdb := rds.NewClient(&rds.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPass,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("ping database redis failed, err: %v", err)
		return nil, err
	}

	return rdb, nil
}
