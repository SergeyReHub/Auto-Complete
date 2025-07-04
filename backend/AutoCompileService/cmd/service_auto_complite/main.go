package main

import (
	"auto_complite/internal/config"
	auto_complete_repository "auto_complite/internal/repository"
	"auto_complite/internal/server"
	"auto_complite/internal/usecase/auto_complete"
	"auto_complite/pkg/graceful_stopping"
	"auto_complite/pkg/logging_interceptor"
	"context"

	"go.uber.org/zap"
)

func main() {
	cfg := config.GetConfig()
	logger := logging_interceptor.InitLogger(cfg)

	repo, err := auto_complete_repository.NewRepository(context.Background(), cfg, logger)
	if err != nil {
		logger.Error("Error create repo.", zap.Error(err))
		return
	}

	usecase := auto_complete.GetAutoCompleteUsecase(repo, logger)

	server, err := server.RegisterAutoCompleteServiceServer(usecase, repo, logger, *cfg)
	if err != nil {
		logger.Error("Error creating server.", zap.Error(err))
		return
	}

	graceful_stopping.GracefulShutDown(server)
}
