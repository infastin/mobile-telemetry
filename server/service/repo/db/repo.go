package db

import "context"

type AtomicCallback func(Repo) error

type Repo interface {
	Update(context.Context, AtomicCallback) error
	View(context.Context, AtomicCallback) error
	Batch(context.Context, AtomicCallback) error
	UserRepo() UserRepo
	DeviceRepo() DeviceRepo
	TelemetryRepo() TelemetryRepo
}
