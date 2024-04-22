package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	lg     *zap.Logger
	server *http.Server
}

type ServerParams struct {
	fx.In

	Logger           *zap.Logger
	TrackHandler     *TrackHandler
	LoggerMiddleware *LoggerMiddleware

	Port int `name:"http_port"`
}

func New(lc fx.Lifecycle, shutdowner fx.Shutdowner, params ServerParams) *Server {
	e := echo.New()

	e.Use(params.LoggerMiddleware.Handle)
	e.Use(NewRecoverMiddleware)

	e.POST("/track", params.TrackHandler.Handle)

	server := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(params.Port),
		Handler: e,
	}

	srv := &Server{
		lg:     params.Logger,
		server: server,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				srv.lg.Info("starting http server")
				if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					srv.lg.Error("could not start http server", zap.Error(err))
					_ = shutdowner.Shutdown()
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.lg.Info("shutting down http server")
			return srv.Shutdown()
		},
	})

	return srv
}

func (s *Server) Run() error {
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("http: failed to server: %w", err)
	}
	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("http: failed to shutdown server: %w", err)
	}
	return nil
}
