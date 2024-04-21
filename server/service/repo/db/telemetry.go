package db

import (
	"context"
	"mobile-telemetry/server/service/model"
)

type TelemetryRepo interface {
	AddTelemetries(ctx context.Context, telemetries []model.Telemetry) (err error)
}
