package schema

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
)

const TelemetryPrefix string = "telemetry"

type Telemetry struct {
	UserID     uuid.UUID
	DeviceID   uint64
	OSVersion  string
	AppVersion string
	Action     string
	Data       map[string]any
	Timestamp  time.Time
}

func TelemetryKey(userID uuid.UUID, deviceID uint64) []byte {
	var b bytes.Buffer

	b.Grow(len(TelemetryPrefix) + 1 + 8)
	b.WriteString(TelemetryPrefix)
	b.WriteByte(':')
	_ = binary.Write(&b, binary.BigEndian, userID)
	b.WriteByte(':')
	_ = binary.Write(&b, binary.BigEndian, deviceID)

	return b.Bytes()
}

func MarshalTelemetryData(data *TelemetryData) ([]byte, error) {
	return data.MarshalMsg(nil)
}

func TelemetryEntry(telemetry *Telemetry) (*badger.Entry, error) {
	key := TelemetryKey(telemetry.UserID, telemetry.DeviceID)

	val, err := MarshalTelemetryData(&TelemetryData{
		OSVersion:  telemetry.OSVersion,
		AppVersion: telemetry.AppVersion,
		Action:     telemetry.Action,
		Data:       telemetry.Data,
		Timestamp:  telemetry.Timestamp,
	})

	if err != nil {
		return nil, err
	}

	return badger.NewEntry(key, val), nil
}