package app

import (
	"mobile-telemetry/server/controller/http"
	trackService "mobile-telemetry/server/service/impl/track"
	dbRepo "mobile-telemetry/server/service/repo/db/entimpl"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func Run() {
	fx.New(
		fx.Provide(NewLogger),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Provide(NewCLI),
		fx.Provide(
			fx.Annotate(
				NewConfig,
				fx.ParamTags(`name:"config_path"`),
			),
		),
		fx.Provide(
			dbRepo.New,
			trackService.New,
		),
		fx.Provide(
			http.NewTrackHandler,
			http.NewLoggerMiddleware,
			http.New,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
