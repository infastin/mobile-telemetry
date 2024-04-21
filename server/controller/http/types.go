package http

import (
	"time"

	"github.com/google/uuid"
)

type GeneralInfo struct {
	UserID     uuid.UUID `json:"user_id" validate:"required"`
	Device     Device    `json:"device" validate:"required"`
	AppVersion string    `json:"app_version" validate:"required"`
}

type Device struct {
	Manufacturer string `json:"manufacturer" validate:"required"`
	Model        string `json:"model" validate:"required"`
	BuildNumber  string `json:"build_number" validate:"required"`
	OS           string `json:"os" validate:"required"`
	OSVersion    string `json:"os_version" validate:"required"`
	ScreenWidth  uint32 `json:"screen_width" validate:"required"`
	ScreenHeight uint32 `json:"screen_height" validate:"required"`
}

type Telemetry struct {
	Action    string         `json:"action" validate:"required"`
	Data      map[string]any `json:"data" validate:"required"`
	Timestamp time.Time      `json:"timestamp" validate:"required"`
}
