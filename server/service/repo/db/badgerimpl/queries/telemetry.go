package queries

import (
	"encoding/binary"
	"mobile-telemetry/pkg/fastconv"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
)

//go:generate msgp -tests=false
//msgp:replace uuid.UUID with:[16]byte
//msgp:ignore TelemetryKey

const TelemetryPrefix = "telemetry"

type TelemetryKey struct {
	ID        uint64
	cachedKey []byte
}

func NewTelemetryKey(id uint64) *TelemetryKey {
	return &TelemetryKey{
		ID:        id,
		cachedKey: nil,
	}
}

func (t *TelemetryKey) Equal(other *TelemetryKey) bool {
	return t.ID == other.ID
}

func (t *TelemetryKey) MarshalBinary() (data []byte, err error) {
	if t.cachedKey != nil {
		return t.cachedKey, nil
	}

	data = append(data, fastconv.Bytes(TelemetryPrefix)...)
	data = append(data, ':')
	data = binary.BigEndian.AppendUint64(data, t.ID)

	t.cachedKey = data

	return t.cachedKey, nil
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

	keyb, _ := NewTelemetryKey(id).MarshalBinary()
	valb, _ := val.MarshalMsg(nil)

	err = tx.Set(keyb, valb)
	if err != nil {
		return 0, err
	}

	return id, nil
}
