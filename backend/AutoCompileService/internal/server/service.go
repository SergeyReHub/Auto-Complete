package server

import (
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
func (a *AutoCompleteServiceServer) AutoComplete(ctx context.Context, req *api.AutoCompleteRequest) (*api.AutoCompleteResponse, error) {
	ac, err := a.us.AutoCompleteText(ctx, req.OriginalString)
	if err != nil {
		a.logg.Error("Service error by autocompleting text", zap.Error(err))
		return nil, err
	}
	return &api.AutoCompleteResponse{
		PrepositionsStrings: ac.PrepositionsStrings,
	}, nil
}
