package model

import (
	"time"

	"github.com/google/uuid"
)

type Telemetry struct {
	UserID     uuid.UUID
	DeviceID   int
	OSVersion  string
	AppVersion string
	Action     string
	Data       map[string]any
	Timestamp  time.Time
}
