package app

const (
	ReleaseMode = "release"
	DebugMode   = "debug"
	TestMode    = "test"
)

var appMode string

func SetMode(mode string) {
	if ValidMode(mode) {
		appMode = mode
	} else {
		panic(`invalid application mode "` + mode + `"`)
	}
}

func Mode() string {
	return appMode
}

func ValidMode(mode string) bool {
	switch mode {
	case ReleaseMode, DebugMode, TestMode:
		return true
	default:
		return false
	}
}
