package http

import (
	"mobile-telemetry/server/entity"
	"mobile-telemetry/server/service"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type TrackHandler struct {
	lg           *zap.Logger
	trackService service.TrackService
}

type TrackHandlerParams struct {
	fx.In

	Logger       *zap.Logger
	TrackService service.TrackService
}

func NewTrackHandler(params TrackHandlerParams) *TrackHandler {
	return &TrackHandler{
		lg:           params.Logger,
		trackService: params.TrackService,
	}
}

type TrackRequest struct {
	Info GeneralInfo `json:"info"`
	Data []Telemetry `json:"data"`
}

func (tr TrackRequest) Validate() error {
	return validation.ValidateStruct(&tr,
		validation.Field(&tr.Info),
		validation.Field(&tr.Data, validation.NotNil),
	)
}

func (h *TrackHandler) Handle(ctx echo.Context) error {
	var req TrackRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user := &entity.User{
		ID: req.Info.UserID,
	}

	device := &entity.Device{
		Manufacturer: req.Info.Device.Manufacturer,
		Model:        req.Info.Device.Model,
		BuildNumber:  req.Info.Device.BuildNumber,
		OS:           req.Info.Device.OS,
		OSVersion:    req.Info.Device.OSVersion,
		ScreenWidth:  req.Info.Device.ScreenWidth,
		ScreenHeight: req.Info.Device.ScreenHeight,
	}

	var data []entity.Telemetry
	for i := 0; i < len(req.Data); i++ {
		data = append(data, entity.Telemetry{
			Action:    req.Data[i].Action,
			Data:      req.Data[i].Data,
			Timestamp: req.Data[i].Timestamp,
		})
	}

	err := h.trackService.AddTelemetry(ctx.Request().Context(),
		service.AddTelemetryParams{
			User:       user,
			Device:     device,
			AppVersion: req.Info.AppVersion,
			Data:       data,
		})
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
