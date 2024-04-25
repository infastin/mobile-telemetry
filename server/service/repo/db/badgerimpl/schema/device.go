package schema

import (
	"bytes"
	"encoding/binary"

	"github.com/dgraph-io/badger/v4"
	"github.com/vmihailenco/msgpack/v5"
)

const DevicePrefix string = "device"

type Device struct {
	ID           uint64
	Manufacturer string
	Model        string
	BuildNumber  string
	OS           string
	ScreenWidth  uint32
	ScreenHeight uint32
}

func DeviceKey(id uint64) []byte {
	var b bytes.Buffer

	b.Grow(len(DevicePrefix) + 1 + 8)
	b.WriteString(DevicePrefix)
	b.WriteByte(':')
	_ = binary.Write(&b, binary.BigEndian, id)

	return b.Bytes()
}

type DeviceData struct {
	Manufacturer string
	Model        string
	BuildNumber  string
	OS           string
	ScreenWidth  uint32
	ScreenHeight uint32
}

func MarshalDeviceData(data *DeviceData) ([]byte, error) {
	var b bytes.Buffer
	err := msgpack.NewEncoder(&b).Encode(data)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func DeviceEntry(device *Device) (*badger.Entry, error) {
	key := DeviceKey(device.ID)

	val, err := MarshalDeviceData(&DeviceData{
		Manufacturer: device.Manufacturer,
		Model:        device.Model,
		BuildNumber:  device.BuildNumber,
		OS:           device.OS,
		ScreenWidth:  device.ScreenWidth,
		ScreenHeight: device.ScreenHeight,
	})

	if err != nil {
		return nil, err
	}

	return badger.NewEntry(key, val), nil
}
