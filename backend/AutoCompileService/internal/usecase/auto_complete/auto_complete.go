package auto_complete

import auto_complete_repository "auto_complite/internal/repository"

type AutoCompleteUsecaseInterface interface {
	AutoCompleteText(originStr string) (string, error)
}

type AutoCompleteUsecase struct {
	autoCompleteRepo auto_complete_repository.Repository
}
