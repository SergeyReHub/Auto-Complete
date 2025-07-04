package auto_complete

import (
	"auto_complite/internal/models"
	auto_complete_repository "auto_complite/internal/repository"
	"context"
	"sync"

	"go.uber.org/zap"
)

type AutoCompleteUsecaseInterface interface {
	AutoCompleteText(ctx context.Context, originStr string) (*models.AutoComplete, error)
}

type AutoCompleteUsecase struct {
	autoCompleteRepo auto_complete_repository.RepositoryUsecase
	logg             *zap.Logger
	mu               sync.Mutex
}

func GetAutoCompleteUsecase(repo auto_complete_repository.RepositoryUsecase, logger *zap.Logger) AutoCompleteUsecaseInterface {
	return &AutoCompleteUsecase{
		autoCompleteRepo: repo,
		logg:             logger,
	}
}

func (us *AutoCompleteUsecase) AutoCompleteText(ctx context.Context, originStr string) (*models.AutoComplete, error) {
	slice, err := us.autoCompleteRepo.FindSimilar(ctx, originStr)
	if err != nil{
		us.logg.Error("Error find similar." + err.Error())
		return nil, err
	}
	return &models.AutoComplete{
		PrepositionsStrings: slice,
	}, nil
}
