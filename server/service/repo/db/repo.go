package db

import "context"

type AtomicCallback func(Repo) error

type Repo interface {
	Atomic(context.Context, AtomicCallback) error
	User() UserRepo
	Device() DeviceRepo
	Telemetry() TelemetryRepo
}
