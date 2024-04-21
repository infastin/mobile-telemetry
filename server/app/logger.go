package app

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerParams struct {
	fx.In

	DebugLevel   zapcore.Level `name:"log_level_debug"`
	ReleaseLevel zapcore.Level `name:"log_level_release"`
}

func NewLogger(params LoggerParams) *zap.Logger {
	var (
		encoder zapcore.Encoder
		level   zapcore.Level
	)

	switch GetMode() {
	case DebugMode:
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		level = params.DebugLevel
	case ReleaseMode:
		encoderConfig := zap.NewProductionEncoderConfig()
		encoder = zapcore.NewJSONEncoder(encoderConfig)
		level = params.ReleaseLevel
	}

	writeSyncer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(encoder, writeSyncer, level)

	return zap.New(core)
}
