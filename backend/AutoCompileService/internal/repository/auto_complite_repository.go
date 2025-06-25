package auto_complete_repository

import (
	"auto_complite/internal/config"
	"auto_complite/internal/repository/redis_storage"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	RedisStorage        *redis_storage.RedisStorage
	PostgresStoragePool *pgxpool.Pool
}

func (*Repository) FindSimilar(ctx context.Context, originStr string) ([]string, error) {
	return nil, nil
}

func NewRepository(cfg *config.Config, logger *zap.Logger) RepositoryUsecase {
	// storage, err := redis_storage.RedisInit(cfg, logger)
	// if err != nil {
	// 	logger.Panic("redis init failed", zap.Error(err))
	// }
	// pool, err := pool_conns.CreatePool(&cfg.AutoFillDB)

	return &Repository{
		RedisStorage:        nil,
		PostgresStoragePool: nil,
	}
}
