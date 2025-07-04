package postgres_storage

import "context"

type PostgresUsecase interface {
	GetSimilarsFromDb(ctx context.Context, str string) ([]string, error)
}
