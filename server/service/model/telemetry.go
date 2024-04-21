package model

import (
	"time"

	"github.com/google/uuid"
)

type Telemetry struct {
	UserUID    uuid.UUID
	DeviceID   int
	OSVersion  string
	AppVersion string
	Action     string
	Data       map[string]any
	Timestamp  time.Time
}
