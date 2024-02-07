package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger

func InitSugar() error {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zapConfig.DisableStacktrace = true
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel) // comment for production
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("Failed to build logger: %s", err.Error())
		return err
	}
	defer logger.Sync() // flushes buffer, if any
	Sugar = logger.Sugar()
	Sugar.Debug("Sugared logger is setuped!")
	return nil
}
