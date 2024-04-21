package entity

type Device struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	BuildNumber  string `json:"build_number"`
	OS           string `json:"os"`
	OSVersion    string `json:"os_version"`
	ScreenWidth  uint32 `json:"screen_width"`
	ScreenHeight uint32 `json:"screen_height"`
}
