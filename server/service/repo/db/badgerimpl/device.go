package impl

import (
	"context"
	"errors"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/badgerimpl/schema"

	"github.com/dgraph-io/badger/v4"
)

func (db *dbRepo) AddDeviceIfNotExists(ctx context.Context, device *model.Device) (id int, err error) {
	devIDKey := schema.DeviceIDKey(device.Manufacturer, device.Model, device.BuildNumber)

	tx := db.db.NewTransaction(true)
	defer tx.Discard()

	devIDItem, err := tx.Get(devIDKey)
	if err != nil && !errors.Is(err, badger.ErrKeyNotFound) {
		return 0, err
	}

	var devID uint64
	if errors.Is(err, badger.ErrKeyNotFound) {
		devID, err = db.deviceSeq.Next()
		if err != nil {
			return 0, err
		}

		devIDVal, err := schema.MarshalDeviceIDData(&schema.DeviceIDData{
			ID: devID,
		})
		if err != nil {
			return 0, err
		}

		err = tx.Set(devIDKey, devIDVal)
		if err != nil {
			return 0, err
		}

		devEntry, err := schema.DeviceEntry(&schema.Device{
			ID:           devID,
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

		err = tx.SetEntry(devEntry)
		if err != nil {
			return 0, err
		}
	} else {
		err = devIDItem.Value(func(val []byte) error {
			devIDData, err := schema.UnmarshalDeviceIDData(val)
			if err != nil {
				return err
			}

			devID = devIDData.ID
			return nil
		})
		if err != nil {
			return 0, err
		}
	}

	userDeviceEntry, err := schema.UserDeviceEntry(&schema.UserDevice{
		UserID:   device.UserID,
		DeviceID: devID,
	})
	if err != nil {
		return 0, err
	}

	err = tx.SetEntry(userDeviceEntry)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(devID), nil
}
