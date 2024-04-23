package main

import (
	"mobile-telemetry/server/app"
	"mobile-telemetry/server/controller/http"
	trackService "mobile-telemetry/server/service/impl/track"
	dbRepo "mobile-telemetry/server/service/repo/db/sqlcimpl"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(app.NewCLI),
		fx.Provide(
			fx.Annotate(
				app.NewConfig,
				fx.ParamTags(`name:"config_path"`),
			),
		),
		fx.Provide(app.NewLogger),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
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
