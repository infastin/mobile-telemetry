package entity

import "time"

type Telemetry struct {
	Action    string
	Data      map[string]any
	Timestamp time.Time
}
