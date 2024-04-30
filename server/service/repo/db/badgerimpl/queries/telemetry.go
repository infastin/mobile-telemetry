package queries

import (
	"encoding/binary"
	"mobile-telemetry/pkg/fastconv"
	"slices"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
)

//go:generate msgp -tests=false
//msgp:replace uuid.UUID with:[16]byte
//msgp:ignore TelemetryKey

const TelemetryPrefix = "telemetry"

type TelemetryKey struct {
	ID uint64
}

func NewTelemetryKey(id uint64) *TelemetryKey {
	return &TelemetryKey{
		ID: id,
	}
}

func (t *TelemetryKey) MarshalKey(b []byte) []byte {
	b = slices.Grow(b, len(TelemetryPrefix)+1+8)
	b = append(b, fastconv.Bytes(TelemetryPrefix)...)
	b = append(b, ':')
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

func (tx *UpdateTx) InsertTelemetry(val *TelemetryValueV1) (id uint64, err error) {
	return insertTelemetry(tx, tx.queries.telemetrySeq, val)
}

func (tx *BatchWriteTx) InsertTelemetry(val *TelemetryValueV1) (id uint64, err error) {
	return insertTelemetry(tx, tx.queries.telemetrySeq, val)
}

func insertTelemetry(tx writeTx, seq *badger.Sequence, val *TelemetryValueV1) (id uint64, err error) {
	id, err = seq.Next()
	if err != nil {
		return 0, err
	}

	keyb := NewTelemetryKey(id).MarshalKey(nil)
	valb, _ := val.MarshalMsg(nil)

	err = tx.Set(keyb, valb)
	if err != nil {
		return 0, err
	}

	return id, nil
}
