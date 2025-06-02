package server

import (
	"auto_complite/internal/usecase/auto_complete"
	"auto_complite/pkg/logging_interceptor"
	api "auto_complite/pkg/proto"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type AutoCompleteServiceServer struct {
	autoCompleteUsecase auto_complete.AutoCompleteUsecaseInterface
	api.UnimplementedAutoCompleteServiceServer
}

func RegisterAutoCompleteServiceServer(log *zap.Logger, cfg config.Config) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logging_interceptor.LoggingInterceptor(log)),
	)

	lis, err := net.Listen("tcp", cfg.GRPCManagement.GetAddr())
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	api.RegisterAutoCompleteServiceServer(grpcServer, NewQuizAutoCompleteServiceServer(quizStorage, clientStorage, redisStorage, log, kfk))

	reflection.Register(grpcServer)

	go func() {
		log.Info("Starting gRPC server", zap.String("port", cfg.GRPCManagement.GRPCPort))
		if err := grpcServer.Serve(lis); err != nil {
			log.Error(fmt.Sprintf("failed to serve: %v", err))
		}
	}()
	return grpcServer, nil
}
