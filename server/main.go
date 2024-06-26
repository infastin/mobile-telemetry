package main

import (
	fxzerolog "mobile-telemetry/pkg/fx-zerolog"
	"mobile-telemetry/server/app"
	"mobile-telemetry/server/controller/http"
	trackService "mobile-telemetry/server/service/impl/track"
	dbRepo "mobile-telemetry/server/service/repo/db/bboltimpl"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
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
		fx.WithLogger(func(logger zerolog.Logger) fxevent.Logger {
			return fxzerolog.New(logger.With().Str("tag", "fx").Logger())
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
