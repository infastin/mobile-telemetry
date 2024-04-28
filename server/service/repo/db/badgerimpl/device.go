package impl

import (
	"context"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/badgerimpl/queries"

	"github.com/dgraph-io/badger/v4"
)

func (db *dbRepo) AddDeviceIfNotExists(ctx context.Context, device *model.Device) (id int, err error) {
	tx := db.queries.Update()
	defer tx.Discard()

	devIdxKey := queries.NewDeviceIndexKey(device.Manufacturer, device.Model, device.BuildNumber)

	devID, err := tx.GetDeviceIndex(devIdxKey)
	if err != nil && err != badger.ErrKeyNotFound {
		return 0, err
	}

	if err == badger.ErrKeyNotFound {
		devID, err = tx.InsertDeviceIndex(devIdxKey)
		if err != nil {
			return 0, err
		}

		err = tx.InsertDevice(queries.NewDeviceKey(devID), &queries.DeviceValueV1{
			Manufacturer: device.Manufacturer,
			Model:        device.Model,
			BuildNumber:  device.BuildNumber,
			OS:           device.OS,
			ScreenWidth:  device.ScreenWidth,
			ScreenHeight: device.ScreenHeight,
		})
		if err != nil {
			return 0, err
		}
	}

	err = tx.InsertUserDevice(queries.NewUserDeviceKey(device.UserID, devID))
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(devID), nil
}
