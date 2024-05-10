package queries

import (
	"encoding/binary"
	"slices"

	"github.com/google/uuid"
)

var UserDeviceBucketName = []byte("user_device")

type UserDeviceKey struct {
	UserID   uuid.UUID
	DeviceID uint64
}

func NewUserDeviceKey(userID uuid.UUID, deviceID uint64) *UserDeviceKey {
	return &UserDeviceKey{
		UserID:   userID,
		DeviceID: deviceID,
	}
}

func (ud *UserDeviceKey) MarshalKey(b []byte) []byte {
	b = slices.Grow(b, 16+8)
	b = append(b, ud.UserID[:]...)
	b = binary.BigEndian.AppendUint64(b, ud.DeviceID)
	return b
}

func (queries *Queries) InsertUserDevice(key *UserDeviceKey) (err error) {
	b := queries.tx.Bucket(UserDeviceBucketName)

	keyb := key.MarshalKey(nil)
	valb := Meta(0).Append(nil)

	return b.Put(keyb, valb)
}
