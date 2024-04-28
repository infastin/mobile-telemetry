package app

import (
	"io"
	"os"
	"path"

	"go.uber.org/fx"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/rs/zerolog"
)

type LoggerParams struct {
	fx.In

	Level      zerolog.Level `name:"logger_level"`
	Directory  string        `name:"logger_directory"`
	MaxSize    int           `name:"logger_max_size"`
	MaxAge     int           `name:"logger_max_age"`
	MaxBackups int           `name:"logger_max_backups"`
}

func NewLogger(params LoggerParams) zerolog.Logger {
	var stdout io.Writer = os.Stdout
	if Mode() != ReleaseMode {
		stdout = zerolog.ConsoleWriter{Out: stdout}
	}

	files := &lumberjack.Logger{
		Filename:   path.Join(params.Directory, "server.log"),
		MaxSize:    params.MaxSize,
		MaxAge:     params.MaxAge,
		MaxBackups: params.MaxBackups,
		LocalTime:  false,
		Compress:   false,
	}

	w := zerolog.MultiLevelWriter(stdout, files)
	lg := zerolog.New(w).Level(params.Level)
	lg = lg.With().Timestamp().Logger()

	return lg
}
