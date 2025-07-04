package redis_storage

import (
	"auto_complite/internal/config"
	"context"
	"errors"
	"strings"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisStorage struct {
	Client *redis.Client
	Logger *zap.Logger
}

func RedisInit(ctx context.Context, cfg *config.Config, logger *zap.Logger) (RedisUsecase, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, errors.New("DB Redis error. Init Redis error.\n" + err.Error())
	}
	return &RedisStorage{
		Client: redisClient, // Просто сохраняем *redis.Client
		Logger: logger,
	}, nil
}

func (s *RedisStorage) SetNewCasheSlice(ctx context.Context, sliceInfoToCache []string) error {
	sliceInterfaces := make([]interface{}, len(sliceInfoToCache))
	for i, v := range sliceInfoToCache {
		sliceInterfaces[i] = v
	}

	err := s.Client.FlushDB(ctx).Err()
	if err != nil {
		s.Logger.Error("DB Redis error. Set value error.", zap.Strings("slice", sliceInfoToCache))
		return errors.New("DB Redis error. Set value error.\n" + err.Error())
	}

	err = s.Client.RPush(ctx, "phrases", sliceInterfaces...).Err()
	if err != nil {
		s.Logger.Error("DB Redis error. Push value error.", zap.Error(err))
		return errors.New("DB Redis error. Push value error.\n" + err.Error())
	}

	return nil
}

func (s *RedisStorage) GetSimilarsFromCache(ctx context.Context, str string) ([]string, error) {
	slice, err := s.Client.LRange(ctx, "phrases", 0, -1).Result()
	if err != nil {
		s.Logger.Error("DB Redis error. Get similars error.", zap.Strings("key", slice))
		return nil, errors.New("DB Redis error. Get similars error.\n" + err.Error())
	}
	var sliceContainsSubstr []string
	for _, v := range slice{
		if strings.Contains(v, str){
			sliceContainsSubstr = append(sliceContainsSubstr, v)
		}
	}
	s.SetNewCasheSlice(ctx, sliceContainsSubstr)
	return slice, nil
}
