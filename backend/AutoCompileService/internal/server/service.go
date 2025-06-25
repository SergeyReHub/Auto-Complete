package server

import (
	"auto_complite/internal/models"
	"auto_complite/internal/usecase/auto_complete"
	api "auto_complite/pkg/proto"
	"context"

	"go.uber.org/zap"
)

type AutoCompleteServiceServer struct {
	api.UnimplementedAutoCompleteServiceServer
	logg *zap.Logger
	us   auto_complete.AutoCompleteUsecaseInterface
}

func NewAutoCompleteServiceServer(logg *zap.Logger, usecase auto_complete.AutoCompleteUsecaseInterface) *AutoCompleteServiceServer {
	return &AutoCompleteServiceServer{
		logg: logg,
		us:   usecase,
	}
}
func (a *AutoCompleteServiceServer) AutoCompleteText(ctx context.Context, originStr string) (*models.AutoComplete, error) {
	return nil, nil
}
