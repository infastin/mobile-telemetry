package impl

import (
	"context"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/sqliteimpl/sqlc"
)

func (db *dbRepo) AddDeviceIfNotExists(ctx context.Context, device *model.Device) (id int, err error) {
	if _, err = db.queries.UpsertDevice(ctx, sqlc.UpsertDeviceParams{
		Manufacturer: device.Manufacturer,
		Model:        device.Model,
		BuildNumber:  device.BuildNumber,
		Os:           device.OS,
		ScreenWidth:  int64(device.ScreenWidth),
		ScreenHeight: int64(device.ScreenHeight),
	}); err != nil {
		return 0, err
	}

	var dev sqlc.Device
	if dev, err = db.queries.FindDevice(ctx, sqlc.FindDeviceParams{
		Manufacturer: device.Manufacturer,
		Model:        device.Model,
		BuildNumber:  device.BuildNumber,
	}); err != nil {
		return 0, err
	}

	if _, err = db.queries.InsertUserDeviceIfNotExists(ctx, sqlc.InsertUserDeviceIfNotExistsParams{
		UserID:   device.UserID,
		DeviceID: dev.ID,
	}); err != nil {
		return 0, err
	}

	return int(dev.ID), nil
}
