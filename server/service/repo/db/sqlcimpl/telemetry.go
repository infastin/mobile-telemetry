package impl

import (
	"context"
	"encoding/json"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db/sqlcimpl/sqlc"
)

func (db *dbRepo) AddTelemetries(ctx context.Context, telemetries []model.Telemetry) (err error) {
	var data []sqlc.InsertTelemetriesBulkParams
	for i := 0; i < len(telemetries); i++ {
		jsonb, err := json.Marshal(telemetries[i].Data)
		if err != nil {
			return err
		}

		data = append(data, sqlc.InsertTelemetriesBulkParams{
			UserID:     telemetries[i].UserUID,
			DeviceID:   int64(telemetries[i].DeviceID),
			OsVersion:  telemetries[i].OSVersion,
			AppVersion: telemetries[i].AppVersion,
			Action:     telemetries[i].Action,
			Data:       jsonb,
			Timestamp:  telemetries[i].Timestamp,
		})
	}

	_, err = db.queries.InsertTelemetriesBulk(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
