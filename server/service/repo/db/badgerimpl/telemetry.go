package impl

import (
	"context"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/badgerimpl/schema"
)

func (db *dbRepo) AddTelemetries(ctx context.Context, telemetries []model.Telemetry) (err error) {
	batch := db.db.NewWriteBatch()
	defer batch.Cancel()

	for i := 0; i < len(telemetries); i++ {
		entry, err := schema.TelemetryEntry(&schema.Telemetry{
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

		err = batch.SetEntry(entry)
		if err != nil {
			return err
		}
	}

	err = batch.Flush()
	if err != nil {
		return err
	}

	return nil
}
