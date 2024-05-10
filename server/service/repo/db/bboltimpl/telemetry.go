package impl

import (
	"context"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/bboltimpl/queries"
)

func (db *dbRepo) AddTelemetries(ctx context.Context, telemetries []model.Telemetry) (err error) {
	for i := 0; i < len(telemetries); i++ {
		if _, err = db.queries.InsertTelemetry(&queries.TelemetryValueV1{
			UserID:     telemetries[i].UserID,
			DeviceID:   uint64(telemetries[i].DeviceID),
			OSVersion:  telemetries[i].OSVersion,
			AppVersion: telemetries[i].AppVersion,
			Action:     telemetries[i].Action,
			Data:       telemetries[i].Data,
			Timestamp:  telemetries[i].Timestamp,
		}); err != nil {
			return err
		}
	}
	return nil
}
