package model

import "github.com/google/uuid"

type Device struct {
	UserID       uuid.UUID
	Manufacturer string
	Model        string
	BuildNumber  string
	OS           string
	ScreenWidth  uint32
	ScreenHeight uint32
}
