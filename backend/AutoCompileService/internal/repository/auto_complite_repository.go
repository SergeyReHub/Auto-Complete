package auto_complete_repository

import (
	"auto_complite/internal/repository/redis_storage"
	"context"
	"go.uber.org/zap"
)

type Repository struct {
	Storage *redis_storage.RedisStorage
}

func (*Repository) FindSimilar(ctx context.Context, originStr string) ([]string, error) {
	return nil, nil
}

func NewRepository(cfg config.Config, logger *zap.Logger) RepositoryUsecase {
	storage, err := redis_storage.RedisInit(cfg, logger)
	if err != nil {
		logger.Panic("redis init failed", zap.Error(err))
	}
	return &Repository{
		Storage: storage,
	}
}
