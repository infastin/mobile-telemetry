package service

import (
	"context"
	"mobile-telemetry/server/entity"
)

type AddTelemetryParams struct {
	User       *entity.User
	Device     *entity.Device
	AppVersion string
	Data       []entity.Telemetry
}

type TrackService interface {
	AddTelemetry(ctx context.Context, params AddTelemetryParams) (err error)
}
