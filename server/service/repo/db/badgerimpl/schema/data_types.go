package schema

import "time"

//go:generate msgp

type DeviceData struct {
	Manufacturer string `msg:"manufacturer"`
	Model        string `msg:"model"`
	BuildNumber  string `msg:"build_number"`
	OS           string `msg:"os"`
	ScreenWidth  uint32 `msg:"screen_width"`
	ScreenHeight uint32 `msg:"screen_height"`
}

type DeviceIDData struct {
	ID uint64 `msg:"id"`
}

type TelemetryData struct {
	OSVersion  string                 `msg:"os_version"`
	AppVersion string                 `msg:"app_version"`
	Action     string                 `msg:"action"`
	Data       map[string]interface{} `msg:"data"`
	Timestamp  time.Time              `msg:"timestamp"`
}

type UserData struct{}

type UserDeviceData struct{}
