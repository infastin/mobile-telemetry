package impl

import (
	"context"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/entimpl/ent"
)

func (db *dbRepo) AddTelemetries(ctx context.Context, telemetries []model.Telemetry) (err error) {
	var builders []*ent.TelemetryCreate
	for i := 0; i < len(telemetries); i++ {
		builders = append(builders, db.client.Telemetry.Create().
			SetUserID(telemetries[i].UserUID).
			SetDeviceID(telemetries[i].DeviceID).
			SetOsVersion(telemetries[i].OSVersion).
			SetAppVersion(telemetries[i].AppVersion).
			SetActionType(telemetries[i].Action).
			SetActionData(telemetries[i].Data).
			SetActionAt(telemetries[i].Timestamp))
	}

	return db.client.Telemetry.CreateBulk(builders...).Exec(ctx)
}
