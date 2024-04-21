package db

import (
	"context"
	"mobile-telemetry/server/service/model"
)

type DeviceRepo interface {
	AddDeviceIfNotExists(ctx context.Context, device *model.Device) (id int, err error)
}
