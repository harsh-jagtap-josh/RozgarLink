package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

var appLogger *zap.SugaredLogger

func init() {
	zapLogger := getLogger()
	appLogger = zapLogger.Sugar()
}

func getLogger() (logger *zap.Logger) {

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		DisableStacktrace: true,
	}

	logger, err := config.Build()
	if err != nil {
		fmt.Printf("failed to initialize logger %v", err)
		return nil
	}
	return logger
}

func Errorw(ctx context.Context, message string, args ...interface{}) {
	appLogger.Errorw(message, args...)
}

func Infow(ctx context.Context, message string, args ...interface{}) {
	appLogger.Infow(message, args...)
}

func Degubw(ctx context.Context, message string, args ...interface{}) {
	appLogger.Debugw(message, args...)
}

func Warnw(ctx context.Context, message string, args ...interface{}) {
	appLogger.Warnw(message, args...)
}

func Fatalw(ctx context.Context, message string, args ...interface{}) {
	appLogger.Fatalw(message, args...)
}
