package server

import (
	"auto_complite/internal/config"
	auto_complete_repository "auto_complite/internal/repository"
	"auto_complite/internal/usecase/auto_complete"
	"auto_complite/pkg/logging_interceptor"
	api "auto_complite/pkg/proto"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RegisterAutoCompleteServiceServer(uc auto_complete.AutoCompleteUsecaseInterface, repo auto_complete_repository.RepositoryUsecase, log *zap.Logger, cfg config.Config) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logging_interceptor.LoggingInterceptor(log)),
	)

	lis, err := net.Listen("tcp", cfg.GRPCAutoFill.GetAddr())
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	api.RegisterAutoCompleteServiceServer(grpcServer, NewAutoCompleteServiceServer(log, uc))

	reflection.Register(grpcServer)

	go func() {
		log.Info("Starting gRPC server", zap.String("port", cfg.GRPCAutoFill.GRPCPort))
		if err := grpcServer.Serve(lis); err != nil {
			log.Error(fmt.Sprintf("failed to serve: %v", err))
		}
	}()
	return grpcServer, nil
}
