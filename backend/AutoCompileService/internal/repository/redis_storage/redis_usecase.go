package redis_storage

import "context"

type RedisUsecase interface {
	SetNewCasheSlice(ctx context.Context, sliceInfoToCache []string) error
	GetSimilarsFromCache(ctx context.Context, str string) ([]string, error)
}
