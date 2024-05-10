package impl

import (
	"context"
	"encoding/json"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/sqliteimpl/sqlc"
)

func (db *dbRepo) AddTelemetries(ctx context.Context, telemetries []model.Telemetry) (err error) {
	for i := 0; i < len(telemetries); i++ {
		data, err := json.Marshal(telemetries[i].Data)
		if err != nil {
			return err
		}

		if _, err = db.queries.InsertTelemetry(ctx, sqlc.InsertTelemetryParams{
			UserID:     telemetries[i].UserID,
			DeviceID:   int64(telemetries[i].DeviceID),
			OsVersion:  telemetries[i].OSVersion,
			AppVersion: telemetries[i].AppVersion,
			Action:     telemetries[i].Action,
			Data:       data,
			Timestamp:  telemetries[i].Timestamp,
		}); err != nil {
			return err
		}
	}

	return nil
}
