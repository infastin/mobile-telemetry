package impl

import (
	"context"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/entimpl/ent"
	entdevice "mobile-telemetry/server/service/repo/db/entimpl/ent/device"
	entuser "mobile-telemetry/server/service/repo/db/entimpl/ent/user"
)

func (db *dbRepo) AddDeviceIfNotExists(ctx context.Context, device *model.Device) (id int, err error) {
	dev, err := db.client.Device.Query().
		Where(entdevice.And(
			entdevice.Manufacturer(device.Manufacturer),
			entdevice.Model(device.Model),
			entdevice.BuildNumber(device.BuildNumber),
		)).
		First(ctx)
	if err == nil {
		_, err = dev.QueryUser().
			Where(entuser.ID(device.UserID)).
			First(ctx)
		if err == nil {
			return dev.ID, nil
		}

		if !ent.IsNotFound(err) {
			return 0, err
		}

		err = dev.Update().
			AddUserIDs(device.UserID).
			Exec(ctx)
		if err != nil {
			return 0, err
		}

		return dev.ID, nil
	}

	if !ent.IsNotFound(err) {
		return 0, err
	}

	dev, err = db.client.Device.Create().
		AddUserIDs(device.UserID).
		SetManufacturer(device.Manufacturer).
		SetModel(device.Model).
		SetBuildNumber(device.BuildNumber).
		SetOs(device.OS).
		SetScreenWidth(device.ScreenWidth).
		SetScreenHeight(device.ScreenHeight).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	return dev.ID, nil
}
