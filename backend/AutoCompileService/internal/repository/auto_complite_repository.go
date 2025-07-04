package auto_complete_repository

import (
	"auto_complite/internal/config"
	"auto_complite/internal/repository/postgres_storage"
	"auto_complite/internal/repository/redis_storage"
	"context"
	"errors"

	"go.uber.org/zap"
)

type Repository struct {
	RedisStorage    redis_storage.RedisUsecase
	PostgresStorage postgres_storage.PostgresUsecase
	Logger          *zap.Logger
}

func (r *Repository) FindSimilar(ctx context.Context, originStr string) ([]string, error) {
	slice, err := r.RedisStorage.GetSimilarsFromCache(ctx, originStr)
	if err != nil {
		r.Logger.Error("Find similars error.", zap.Error(err))
		return nil, errors.New("Find similars error." + err.Error())
	}
	if len(slice) == 0 {
		slice, err = r.PostgresStorage.GetSimilarsFromDb(ctx, originStr)
		if err != nil {
			r.Logger.Error("Find similars error.", zap.Error(err))
			return nil, errors.New("Find similars error." + err.Error())
		}
		err = r.RedisStorage.SetNewCasheSlice(ctx, slice)
		if err != nil {
			r.Logger.Error("Find similars error.", zap.Error(err))
			return nil, errors.New("Find similars error." + err.Error())
		}
	}
	return slice, nil
}

func NewRepository(ctx context.Context, cfg *config.Config, logger *zap.Logger) (RepositoryUsecase, error) {
	redis_storage, err := redis_storage.RedisInit(ctx, cfg, logger)
	if err != nil {
		logger.Panic("Redis init failed", zap.Error(err))
		return nil, err
	}
	pg_storage, err := postgres_storage.PostgresInit(&cfg.AutoFillDB, logger)
	if err != nil {
		logger.Panic("Postgres init failed", zap.Error(err))
		return nil, err
	}

	return &Repository{
		RedisStorage:    redis_storage,
		PostgresStorage: pg_storage,
		Logger:          logger,
	}, nil
}
