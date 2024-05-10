package queries

import (
	"encoding/binary"
	"slices"
)

//go:generate msgp -tests=false
//msgp:ignore DeviceKey

var DeviceBucketName = []byte("device")

type DeviceKey struct {
	ID uint64
}

func NewDeviceKey(id uint64) *DeviceKey {
	return &DeviceKey{
		ID: id,
	}
}

func (d *DeviceKey) MarshalKey(b []byte) []byte {
	b = slices.Grow(b, 8)
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

func (queries *Queries) InsertDevice(key *DeviceKey, val *DeviceValueV1) (err error) {
	b := queries.tx.Bucket(DeviceBucketName)

	keyb := key.MarshalKey(nil)

	valb := Meta(0).Append(nil)
	valb, _ = val.MarshalMsg(valb)

	return b.Put(keyb, valb)
}
