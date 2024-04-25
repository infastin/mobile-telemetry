package schema

import (
	"bytes"
	"encoding/binary"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
)

const UserDevicePrefix string = "user_device"

type UserDevice struct {
	UserID   uuid.UUID
	DeviceID uint64
}

func UserDeviceKey(userID uuid.UUID, devID uint64) []byte {
	var b bytes.Buffer

	b.Grow(len(UserDevicePrefix) + 1 + len(userID) + 1 + 8)
	b.WriteString(UserDevicePrefix)
	b.WriteByte(':')
	_ = binary.Write(&b, binary.BigEndian, userID)
	b.WriteByte(':')
	_ = binary.Write(&b, binary.BigEndian, devID)

	return b.Bytes()
}

func MarshalUserDeviceData(data *UserDeviceData) ([]byte, error) {
	return data.MarshalMsg(nil)
}

func UserDeviceEntry(userDevice *UserDevice) (*badger.Entry, error) {
	key := UserDeviceKey(userDevice.UserID, userDevice.DeviceID)

	val, err := MarshalUserDeviceData(&UserDeviceData{})
	if err != nil {
		return nil, err
	}

	return badger.NewEntry(key, val), nil
}
