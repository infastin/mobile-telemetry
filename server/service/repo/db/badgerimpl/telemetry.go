package impl

import (
	"context"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/badgerimpl/queries"
)

func (db *dbRepo) AddTelemetries(ctx context.Context, telemetries []model.Telemetry) (err error) {
	batch := db.queries.BatchWrite()
	defer batch.Discard()

	for i := 0; i < len(telemetries); i++ {
		_, err = batch.InsertTelemetry(&queries.TelemetryValueV1{
			UserID:     telemetries[i].UserUID,
			DeviceID:   uint64(telemetries[i].DeviceID),
			OSVersion:  telemetries[i].OSVersion,
			AppVersion: telemetries[i].AppVersion,
			Action:     telemetries[i].Action,
			Data:       telemetries[i].Data,
			Timestamp:  telemetries[i].Timestamp,
		})
		if err != nil {
			return err
		}
	}

	err = batch.Commit()
	if err != nil {
		return err
	}

	return nil
}
