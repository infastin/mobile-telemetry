package http

import (
	"time"

	"github.com/google/uuid"
	"github.com/infastin/go-validation"
)

type GeneralInfo struct {
	UserID     uuid.UUID `json:"user_id"`
	Device     Device    `json:"device"`
	AppVersion string    `json:"app_version"`
}

func (info *GeneralInfo) Validate() error {
	return validation.All(
		validation.Comparable(info.UserID, "user_id").Required(true),
		validation.Ptr(&info.Device, "device").With(validation.Custom),
		validation.String(info.AppVersion, "app_version").Required(true),
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

func (d *Device) Validate() error {
	return validation.All(
		validation.String(d.Manufacturer, "manufacturer").Required(true),
		validation.String(d.Model, "model").Required(true),
		validation.String(d.BuildNumber, "build_number").Required(true),
		validation.String(d.OS, "os").Required(true),
		validation.Number(d.ScreenWidth, "screen_width").Required(true),
		validation.Number(d.ScreenHeight, "screen_height").Required(true),
	)
}

type Telemetry struct {
	Action    string         `json:"action"`
	Data      map[string]any `json:"data"`
	Timestamp time.Time      `json:"timestamp"`
}

func (t Telemetry) Validate() error {
	return validation.All(
		validation.String(t.Action, "action").Required(true),
		validation.Map(t.Data, "data").NotNil(true),
		validation.Time(t.Timestamp, "timestamp").Required(true),
	)
}
