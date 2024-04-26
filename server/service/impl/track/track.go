package impl

import (
	"context"
	"mobile-telemetry/server/service"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type trackService struct {
	lg *zap.Logger
	db db.Repo
}

type TrackServiceParams struct {
	fx.In

	Logger   *zap.Logger
	Database db.Repo
}

func New(params TrackServiceParams) service.TrackService {
	return &trackService{
		lg: params.Logger,
		db: params.Database,
	}
}

func (ts *trackService) AddTelemetry(ctx context.Context, params service.AddTelemetryParams) (err error) {
	userID, err := ts.db.AddUserIfNotExists(ctx, params.User)
	if err != nil {
		return err
	}

	deviceID, err := ts.db.AddDeviceIfNotExists(ctx,
		&model.Device{
			UserID:       params.User.ID,
			Manufacturer: params.Device.Manufacturer,
			Model:        params.Device.Model,
			BuildNumber:  params.Device.BuildNumber,
			OS:           params.Device.OS,
			ScreenWidth:  params.Device.ScreenWidth,
			ScreenHeight: params.Device.ScreenHeight,
		})
	if err != nil {
		return err
	}

	var telemetries []model.Telemetry
	for i := 0; i < len(params.Data); i++ {
		telemetries = append(telemetries, model.Telemetry{
			UserUID:    userID,
			DeviceID:   deviceID,
			OSVersion:  params.Device.OSVersion,
			AppVersion: params.AppVersion,
			Action:     params.Data[i].Action,
			Data:       params.Data[i].Data,
			Timestamp:  params.Data[i].Timestamp,
		})
	}

	return ts.db.AddTelemetries(ctx, telemetries)
}
