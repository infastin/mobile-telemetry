package app

import (
	"fmt"
	"mobile-telemetry/pkg/fastconv"
)

type Mode string

const (
	ReleaseMode Mode = "release"
	DebugMode   Mode = "debug"
)

func (am *Mode) UnmarshalText(text []byte) error {
	switch fastconv.String(text) {
	case string(ReleaseMode):
		*am = ReleaseMode
	case string(DebugMode):
		*am = DebugMode
	default:
		return fmt.Errorf(
			`invalid application mode, got "%s", expected one of "%s", "%s"`,
			fastconv.String(text),
			ReleaseMode, DebugMode,
		)
	}

	return nil
}

var appMode Mode

func SetMode(mode Mode) {
	appMode = mode
}

func GetMode() Mode {
	return appMode
}
