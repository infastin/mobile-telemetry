package queries

import (
	"encoding/binary"
	"mobile-telemetry/pkg/fastconv"
)

//go:generate msgp -tests=false
//msgp:ignore DeviceKey

const DevicePrefix = "device"

type DeviceKey struct {
	ID        uint64
	cachedKey []byte
}

func NewDeviceKey(id uint64) *DeviceKey {
	return &DeviceKey{
		ID:        id,
		cachedKey: nil,
	}
}

func (d *DeviceKey) Equal(other *DeviceKey) bool {
	return d.ID == other.ID
}

func (d *DeviceKey) MarshalBinary() (data []byte, err error) {
	if d.cachedKey != nil {
		return d.cachedKey, nil
	}

	data = append(data, fastconv.Bytes(DevicePrefix)...)
	data = append(data, ':')
	data = binary.BigEndian.AppendUint64(data, d.ID)

	d.cachedKey = data

	return d.cachedKey, nil
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

func insertDevice(setter Setter, key *DeviceKey, val *DeviceValueV1) (err error) {
	keyb, _ := key.MarshalBinary()
	valb, _ := val.MarshalMsg(nil)
	return setter.Set(keyb, valb)
}
