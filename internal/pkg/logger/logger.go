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
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel), // can set to any level like InfoLevel, ErrorLevel, DebugLevel
		Development:       false,                               // takes the stack traces more, kept false so now it wont
		Encoding:          "console",                           // two options json and console, using console will return it in console
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),   // can be either Development congig or Production config
		OutputPaths:       []string{"stdout"},                  // all the output paths, add output to all the paths specified
		DisableStacktrace: true,
	}

	logger, err := config.Build() // to build a logger based on the config provided
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
