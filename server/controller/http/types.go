package http

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type GeneralInfo struct {
	UserID     uuid.UUID `json:"user_id"`
	Device     Device    `json:"device"`
	AppVersion string    `json:"app_version"`
}

func (info GeneralInfo) Validate() error {
	return validation.ValidateStruct(&info,
		validation.Field(&info.UserID, validation.By(ValidUUID)),
		validation.Field(&info.Device),
		validation.Field(&info.AppVersion, validation.Required),
	)
}

type Device struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	BuildNumber  string `json:"build_number"`
	OS           string `json:"os"`
	OSVersion    string `json:"os_version"`
	ScreenWidth  uint32 `json:"screen_width"`
	ScreenHeight uint32 `json:"screen_height"`
}

func (d Device) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Manufacturer, validation.Required),
		validation.Field(&d.Model, validation.Required),
		validation.Field(&d.BuildNumber, validation.Required),
		validation.Field(&d.OS, validation.Required),
		validation.Field(&d.ScreenWidth, validation.Required),
		validation.Field(&d.ScreenHeight, validation.Required),
	)
}

type Telemetry struct {
	Action    string         `json:"action"`
	Data      map[string]any `json:"data"`
	Timestamp time.Time      `json:"timestamp"`
}

func (t Telemetry) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Action, validation.Required),
		validation.Field(&t.Data, validation.NotNil),
		validation.Field(&t.Timestamp, validation.Required),
	)
}
