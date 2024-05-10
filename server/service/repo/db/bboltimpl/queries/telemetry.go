package queries

import (
	"encoding/binary"
	"slices"
	"time"

	"github.com/google/uuid"
)

//go:generate msgp -tests=false
//msgp:replace uuid.UUID with:[16]byte
//msgp:ignore TelemetryKey

var TelemetryBucketName = []byte("telemetry")

type TelemetryKey struct {
	ID uint64
}

func NewTelemetryKey(id uint64) *TelemetryKey {
	return &TelemetryKey{
		ID: id,
	}
}

func (t *TelemetryKey) MarshalKey(b []byte) []byte {
	b = slices.Grow(b, 8)
	b = binary.BigEndian.AppendUint64(b, t.ID)
	return b
}

type TelemetryValueV1 struct {
	UserID     uuid.UUID      `msg:"user_id"`
	DeviceID   uint64         `msg:"device_id"`
	OSVersion  string         `msg:"os_version"`
	AppVersion string         `msg:"app_version"`
	Action     string         `msg:"action"`
	Data       map[string]any `msg:"data"`
	Timestamp  time.Time      `msg:"timestamp"`
}

func (queries *Queries) InsertTelemetry(val *TelemetryValueV1) (id uint64, err error) {
	b := queries.tx.Bucket(TelemetryBucketName)

	id, err = b.NextSequence()
	if err != nil {
		return 0, err
	}

	keyb := NewTelemetryKey(id).MarshalKey(nil)

	valb := Meta(0).Append(nil)
	valb, _ = val.MarshalMsg(valb)

	err = b.Put(keyb, valb)
	if err != nil {
		return 0, err
	}

	return id, nil
}
