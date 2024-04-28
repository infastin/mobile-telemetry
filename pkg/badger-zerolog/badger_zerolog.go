package badgerzerolog

import (
	"fmt"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func New(lg zerolog.Logger) *Logger {
	return &Logger{
		logger: lg,
	}
}

func (l *Logger) Errorf(format string, args ...any) {
	l.logger.Error().Msg(fmt.Sprintf(format, args...))
}

func (l *Logger) Warningf(format string, args ...any) {
	l.logger.Warn().Msg(fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...any) {
	l.logger.Info().Msg(fmt.Sprintf(format, args...))
}

func (l *Logger) Debugf(format string, args ...any) {
	l.logger.Debug().Msg(fmt.Sprintf(format, args...))
}
