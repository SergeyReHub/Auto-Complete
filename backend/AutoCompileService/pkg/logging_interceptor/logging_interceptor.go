package logging_interceptor

import (
	"auto_complite/internal/config"
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

type contextKey string

const (
	Key       contextKey = "logger"
	RequestID contextKey = "request_id"
)

type Logger struct {
	l *zap.Logger
}

func LoggingInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Info("gRPC request",
			zap.String("method", info.FullMethod),
			zap.Any("request", req),
		)

		resp, err = handler(ctx, req)
		if err != nil {
			log.Error("gRPC error",
				zap.String("method", info.FullMethod),
				zap.Error(err),
			)
		}

		return resp, err
	}
}

func InitLogger(cfg *config.Config) *zap.Logger {
	var logger *zap.Logger
	var err error

	logLevel := zapcore.InfoLevel
	switch cfg.Logger.Level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	}

	if cfg.Logger.Level == "prod" {
		prodConfig := zap.NewProductionConfig()
		prodConfig.Level = zap.NewAtomicLevelAt(logLevel)
		logger, err = prodConfig.Build()
	} else {
		devConfig := zap.NewDevelopmentConfig()
		devConfig.Level = zap.NewAtomicLevelAt(logLevel)
		devConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = devConfig.Build()
	}

	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}

	logger = logger.With(
		zap.String("service", "auto-complete-service"),
	)

	zap.ReplaceGlobals(logger)

	return logger
}

func SyncLogger(logger *zap.Logger) {
	_ = logger.Sync()
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}

	l.l.Info(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}

	l.l.Error(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}

	l.l.Fatal(msg, fields...)
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}

	l.l.Debug(msg, fields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}

	l.l.Warn(msg, fields...)
}

func (l *Logger) Sync() error {
	if err := l.l.Sync(); err != nil {
		return fmt.Errorf("failed to sync logger: %w", err)
	}

	return nil
}
