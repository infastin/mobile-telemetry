package db

type Repo interface {
	UserRepo
	DeviceRepo
	TelemetryRepo
}
