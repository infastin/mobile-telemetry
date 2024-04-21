package http

import (
	"errors"
	"fmt"
	"mobile-telemetry/pkg/fastconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	lg *zap.Logger
}

func NewLoggerMiddleware(lg *zap.Logger) *LoggerMiddleware {
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
		}

		elapsed := time.Since(start)

		lg := m.lg.With(
			zap.String("method", req.Method),
			zap.String("uri", req.RequestURI),
			zap.String("addr", req.RemoteAddr),
			zap.Int("status", resp.Status),
			zap.Duration("elapsed", elapsed),
		)

		switch v := err.(type) {
		case *PanicError:
			lg.Error("request panic",
				zap.Any("panic", v.Panic),
				zap.String("stack", fastconv.String(v.Stack)),
			)
		case error:
			var msg string
			if echoErr := (*echo.HTTPError)(nil); errors.As(err, &echoErr) {
				msg = fmt.Sprint(echoErr.Message)
			} else {
				msg = err.Error()
			}
			lg.Error("request error", zap.String("error", msg))
		case nil:
			lg.Info("request")
		}

		return nil
	}
}
