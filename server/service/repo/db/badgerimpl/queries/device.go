package queries

import (
	"encoding/binary"
	"mobile-telemetry/pkg/fastconv"
	"slices"
)

//go:generate msgp -tests=false
//msgp:ignore DeviceKey

const DevicePrefix = "device"

type DeviceKey struct {
	ID uint64
}

func NewDeviceKey(id uint64) *DeviceKey {
	return &DeviceKey{
		ID: id,
	}
}

func (d *DeviceKey) MarshalKey(b []byte) []byte {
	b = slices.Grow(b, len(DevicePrefix)+1+8)
	b = append(b, fastconv.Bytes(DevicePrefix)...)
	b = append(b, ':')
	b = binary.BigEndian.AppendUint64(b, d.ID)
	return b
}

type DeviceValueV1 struct {
	Manufacturer string `msg:"manufacturer"`
	Model        string `msg:"model"`
	BuildNumber  string `msg:"build_number"`
	OS           string `msg:"os"`
	ScreenWidth  uint32 `msg:"screen_width"`
	ScreenHeight uint32 `msg:"screen_height"`
}

func (tx *UpdateTx) InsertDevice(key *DeviceKey, val *DeviceValueV1) (err error) {
	return insertDevice(tx, key, val)
}

func insertDevice(tx writeTx, key *DeviceKey, val *DeviceValueV1) (err error) {
	keyb := key.MarshalKey(nil)
	valb, _ := val.MarshalMsg(nil)
	return tx.Set(keyb, valb)
}
