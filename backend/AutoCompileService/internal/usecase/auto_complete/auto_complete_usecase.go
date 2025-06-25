package auto_complete

import (
	auto_complete_repository "auto_complite/internal/repository"
	"sync"

	"go.uber.org/zap"
)

type AutoCompleteUsecaseInterface interface {
	AutoCompleteText(originStr string) (string, error)
}

type AutoCompleteUsecase struct {
	autoCompleteRepo auto_complete_repository.RepositoryUsecase
	logg             *zap.Logger
	mu               sync.Mutex
}

func GetAutoCompleteUsecase (repo auto_complete_repository.RepositoryUsecase, logger *zap.Logger) AutoCompleteUsecaseInterface{
	return &AutoCompleteUsecase{
		autoCompleteRepo: repo,
		logg: logger,
	}
}

func (*AutoCompleteUsecase) AutoCompleteText(originStr string) (string, error){
	return "", nil
}
