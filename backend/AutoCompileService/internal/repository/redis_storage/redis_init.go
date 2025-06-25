package redis_storage

import (
	"auto_complite/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisStorage struct {
	Client *redis.Client
	Logger *zap.Logger
}

func RedisInit(cfg *config.Config, logger *zap.Logger) (*RedisStorage, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisStorage{
		Client: redisClient, // Просто сохраняем *redis.Client
		Logger: logger,
	}, nil
}
