package http

import (
	"context"
	"errors"
	"fmt"
	"mobile-telemetry/server/app"
	"net/http"
	"strconv"
	"time"

	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"

	"go.uber.org/fx"
)

type Server struct {
	lg     zerolog.Logger
	server *http.Server
}

type ServerParams struct {
	fx.In

	Logger           zerolog.Logger
	TrackHandler     *TrackHandler
	LoggerMiddleware *LoggerMiddleware

	Port int `name:"http_port"`
}

func New(lc fx.Lifecycle, shutdowner fx.Shutdowner, params ServerParams) *Server {
	e := echo.New()

	e.JSONSerializer = JSONSerializer{}

	e.Use(params.LoggerMiddleware.Handle)
	e.Use(NewRecoverMiddleware)

	e.POST("/track", params.TrackHandler.Handle)
	if app.Mode() == app.DebugMode {
		e.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))
	}

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
				srv.lg.Info().Msg("starting http server")
				if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					srv.lg.Err(err).Msg("could not start http server")
					_ = shutdowner.Shutdown()
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.lg.Info().Msg("shutting down http server")
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
