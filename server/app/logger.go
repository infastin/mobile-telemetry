package app

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	writeSyncer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	return zap.New(core)
}
