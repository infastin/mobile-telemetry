package http

import (
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type LoggerMiddleware struct {
	lg zerolog.Logger
}

func NewLoggerMiddleware(lg zerolog.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		lg: lg,
	}
}

func (m *LoggerMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		resp := c.Response()
		start := time.Now()

		err := next(c)
		if err != nil {
			c.Error(err)
		} else if m.lg.GetLevel() > zerolog.InfoLevel {
			return nil
		}

		elapsed := time.Since(start)

		lgCtx := m.lg.With().
			Str("method", req.Method).
			Str("uri", req.RequestURI).
			Str("addr", req.RemoteAddr).
			Int("status", resp.Status).
			Dur("elapsed", elapsed)

		switch v := err.(type) {
		case *PanicError:
			lgCtx = lgCtx.
				Any("panic", v.Panic).
				Bytes("stack", v.Stack)
			lg := lgCtx.Logger()
			lg.Error().Msg("request panic")
		case error:
			if echoErr := (*echo.HTTPError)(nil); errors.As(err, &echoErr) {
				lgCtx = lgCtx.Any("error", echoErr.Message)
			} else {
				lgCtx = lgCtx.Err(err)
			}
			lg := lgCtx.Logger()
			lg.Error().Msg("request error")
		case nil:
			lg := lgCtx.Logger()
			lg.Info().Msg("request")
		}

		return nil
	}
}
