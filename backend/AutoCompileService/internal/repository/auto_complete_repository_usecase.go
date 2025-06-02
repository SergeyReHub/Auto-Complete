package auto_complete_repository

import (
	"context"
)

type RepositoryUsecase interface {
	FindSimilar(ctx context.Context, originStr string) ([]string, error)
}
