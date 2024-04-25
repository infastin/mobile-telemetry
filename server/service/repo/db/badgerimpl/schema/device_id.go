package schema

import (
	"bytes"

	"github.com/dgraph-io/badger/v4"
	"github.com/vmihailenco/msgpack/v5"
)

const DeviceIDPrefix string = "device_id"

type DeviceID struct {
	Manufacturer string
	Model        string
	BuildNumber  string
	ID           uint64
}

func DeviceIDKey(manufacturer, model, buildNumber string) []byte {
	var b bytes.Buffer

	b.Grow(len(DevicePrefix) + 1 + len(manufacturer) + 1 + len(model) + 1 + len(buildNumber))
	b.WriteString(DevicePrefix)
	b.WriteByte(':')
	b.WriteString(manufacturer)
	b.WriteByte(':')
	b.WriteString(model)
	b.WriteByte(':')
	b.WriteString(buildNumber)

	return b.Bytes()
}

type DeviceIDData struct {
	ID uint64
}

func MarshalDeviceIDData(data *DeviceIDData) ([]byte, error) {
	var b bytes.Buffer
	err := msgpack.NewEncoder(&b).Encode(data)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func UnmarshalDeviceIDData(b []byte) (data DeviceIDData, err error) {
	err = msgpack.Unmarshal(b, &data)
	if err != nil {
		return DeviceIDData{}, err
	}
	return data, nil
}

func DeviceIDEntry(deviceID *DeviceID) (*badger.Entry, error) {
	key := DeviceIDKey(deviceID.Manufacturer, deviceID.Model, deviceID.BuildNumber)

	val, err := MarshalDeviceIDData(&DeviceIDData{
		ID: deviceID.ID,
	})

	if err != nil {
		return nil, err
	}

	return badger.NewEntry(key, val), nil
}
