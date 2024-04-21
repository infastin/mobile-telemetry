package impl

import (
	"context"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/sqlcimpl/sqlc"
)

func (db *dbRepo) AddDeviceIfNotExists(ctx context.Context, device *model.Device) (id int, err error) {
	_, err = db.queries.UpserDevice(ctx,
		sqlc.UpserDeviceParams{
			Manufacturer: device.Manufacturer,
			Model:        device.Model,
			BuildNumber:  device.BuildNumber,
			Os:           device.OS,
			ScreenWidth:  int32(device.ScreenWidth),
			ScreenHeight: int32(device.ScreenHeight),
		})
	if err != nil {
		return 0, err
	}

	dev, err := db.queries.FindDevice(ctx,
		sqlc.FindDeviceParams{
			Manufacturer: device.Manufacturer,
			Model:        device.Model,
			BuildNumber:  device.BuildNumber,
		})
	if err != nil {
		return 0, err
	}

	_, err = db.queries.UpsertUserDevice(ctx,
		sqlc.UpsertUserDeviceParams{
			UserID:   device.UserID,
			DeviceID: dev.ID,
		})
	if err != nil {
		return 0, err
	}

	return int(dev.ID), nil
}
