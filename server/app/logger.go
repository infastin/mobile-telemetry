package app

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerParams struct {
	fx.In

	Level zapcore.Level `name:"logger_level"`
}

func NewLogger(params LoggerParams) *zap.Logger {
	var encoder zapcore.Encoder

	switch Mode() {
	case DebugMode:
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case ReleaseMode:
		encoderConfig := zap.NewProductionEncoderConfig()
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	writeSyncer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(encoder, writeSyncer, params.Level)

	return zap.New(core)
}
