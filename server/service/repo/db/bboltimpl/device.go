package impl

import (
	"context"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/bboltimpl/queries"
)

func (db *dbRepo) AddDeviceIfNotExists(ctx context.Context, device *model.Device) (id int, err error) {
	devIdxKey := queries.NewDeviceIndexKey(device.Manufacturer, device.Model, device.BuildNumber)
	devIdxVal := &queries.DeviceIndexValue{
		Manufacturer: device.Manufacturer,
		Model:        device.Model,
		BuildNumber:  device.BuildNumber,
	}

	devID, err := db.queries.FindDeviceIndex(devIdxKey, devIdxVal)
	if err != nil && err != queries.ErrKeyNotFound {
		return 0, err
	}

	if err == queries.ErrKeyNotFound {
		devID, err = db.queries.InsertDeviceIndex(devIdxKey, devIdxVal)
		if err != nil {
			return 0, err
		}

		if err := db.queries.InsertDevice(queries.NewDeviceKey(devID), &queries.DeviceValueV1{
			Manufacturer: device.Manufacturer,
			Model:        device.Model,
			BuildNumber:  device.BuildNumber,
			OS:           device.OS,
			ScreenWidth:  device.ScreenWidth,
			ScreenHeight: device.ScreenHeight,
		}); err != nil {
			return 0, err
		}
	}

	err = db.queries.InsertUserDevice(queries.NewUserDeviceKey(device.UserID, devID))
	if err != nil {
		return 0, err
	}

	return int(devID), nil
}
