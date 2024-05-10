package impl

import (
	"context"
	"mobile-telemetry/server/entity"
	"mobile-telemetry/server/service"
	"mobile-telemetry/server/service/model"
	"mobile-telemetry/server/service/repo/db"
	"time"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type trackService struct {
	lg zerolog.Logger
	db db.Repo

	telemetryChan nbChan[*telemetryRequest]
	flushChan     nbChan[*flushData]
	quitChan      chan struct{}
	didQuitChan   chan struct{}
}

type TrackServiceParams struct {
	fx.In

	Logger   zerolog.Logger
	Database db.Repo
}

func New(lc fx.Lifecycle, params TrackServiceParams) service.TrackService {
	const numFlushers = 4

	ts := &trackService{
		lg:            params.Logger,
		db:            params.Database,
		telemetryChan: make(nbChan[*telemetryRequest], 1<<18),
		flushChan:     make(nbChan[*flushData], 1<<6),
		quitChan:      make(chan struct{}, 1+numFlushers),
		didQuitChan:   make(chan struct{}, 1+numFlushers),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				ts.processRequests()
			}()
			for range numFlushers {
				go func() {
					ts.processFlush()
				}()
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			for range cap(ts.quitChan) {
				ts.quitChan <- struct{}{}
			}
			for range cap(ts.didQuitChan) {
				<-ts.didQuitChan
			}
			return nil
		},
	})

	return ts
}

type flushData struct {
	users       []entity.User
	devices     []model.Device
	telemetries [][]model.Telemetry
	len         int
}

func newFlushData() *flushData {
	return &flushData{
		users:       make([]entity.User, 0),
		devices:     make([]model.Device, 0),
		telemetries: make([][]model.Telemetry, 0),
		len:         0,
	}
}

func (d *flushData) append(req *telemetryRequest) {
	d.users = append(d.users, req.user)
	d.devices = append(d.devices, req.device)
	d.telemetries = append(d.telemetries, req.telemetries)
	d.len += 1
}

func (d *flushData) reset() *flushData {
	old := new(flushData)
	*old = *d
	*d = *newFlushData()
	return old
}

type telemetryRequest struct {
	user        entity.User
	device      model.Device
	telemetries []model.Telemetry
}

func (ts *trackService) processRequests() {
	ticker := time.NewTicker(time.Hour)
	data := newFlushData()

	for {
		select {
		case r := <-ts.telemetryChan:
			data.append(r)
			if data.len >= 1<<12 {
				if dropped := ts.flushChan.send(data.reset()); dropped {
					ts.lg.Warn().Msg("dropped flush data")
				}
			}
		case <-ticker.C:
			if data.len != 0 {
				if dropped := ts.flushChan.send(data.reset()); dropped {
					ts.lg.Warn().Msg("dropped flush data")
				}
			}
		case <-ts.quitChan:
			if data.len != 0 {
				if dropped := ts.flushChan.send(data.reset()); dropped {
					ts.lg.Warn().Msg("dropped flush data")
				}
			}
			ts.didQuitChan <- struct{}{}
			return
		}
	}
}

func (ts *trackService) processFlush() {
	for {
		select {
		case data := <-ts.flushChan:
			ts.flushTelemetries(data)
		case <-ts.quitChan:
			for {
				select {
				case data := <-ts.flushChan:
					ts.flushTelemetries(data)
				default:
					ts.didQuitChan <- struct{}{}
					return
				}
			}
		}
	}
}

func (ts *trackService) flushTelemetries(data *flushData) {
	ctx := context.Background()

	if err := ts.db.Batch(ctx, func(r db.Repo) error {
		for i := 0; i < len(data.users); i++ {
			userID, err := r.UserRepo().AddUserIfNotExists(ctx, &data.users[i])
			if err != nil {
				return err
			}

			data.devices[i].UserID = userID
			for j := 0; j < len(data.telemetries[i]); j++ {
				data.telemetries[i][j].UserID = userID
			}
		}

		for i := 0; i < len(data.devices); i++ {
			deviceID, err := r.DeviceRepo().AddDeviceIfNotExists(ctx, &data.devices[i])
			if err != nil {
				return err
			}

			for j := 0; j < len(data.telemetries[i]); j++ {
				data.telemetries[i][j].DeviceID = deviceID
			}
		}

		for i := 0; i < len(data.telemetries); i++ {
			err := r.TelemetryRepo().AddTelemetries(ctx, data.telemetries[i])
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		ts.lg.Err(err).Msg("flush telemetries")
		return
	}
}

func (ts *trackService) AddTelemetry(ctx context.Context, params service.AddTelemetryParams) (err error) {
	req := &telemetryRequest{
		user: entity.User{
			ID: params.User.ID,
		},
		device: model.Device{
			Manufacturer: params.Device.Manufacturer,
			Model:        params.Device.Model,
			BuildNumber:  params.Device.BuildNumber,
			OS:           params.Device.OS,
			ScreenWidth:  params.Device.ScreenWidth,
			ScreenHeight: params.Device.ScreenHeight,
		},
		telemetries: make([]model.Telemetry, len(params.Data)),
	}

	for i := 0; i < len(params.Data); i++ {
		req.telemetries[i] = model.Telemetry{
			OSVersion:  params.Device.OSVersion,
			AppVersion: params.AppVersion,
			Action:     params.Data[i].Action,
			Data:       params.Data[i].Data,
			Timestamp:  params.Data[i].Timestamp,
		}
	}

	if dropped := ts.telemetryChan.send(req); dropped {
		ts.lg.Warn().Msg("dropped request")
	}

	return nil
}
