package queries

import (
	"encoding/binary"
	"slices"

	"github.com/google/uuid"
)

const UserDevicePrefix = "user_device"

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
	b = slices.Grow(b, len(UserDevicePrefix)+1+16+1+8)
	b = append(b, UserDevicePrefix...)
	b = append(b, ':')
	b = append(b, ud.UserID[:]...)
	b = append(b, ':')
	b = binary.BigEndian.AppendUint64(b, ud.DeviceID)
	return b
}

func (tx *UpdateTx) InsertUserDevice(key *UserDeviceKey) (err error) {
	return insertUserDevice(tx, key)
}

func insertUserDevice(tx writeTx, key *UserDeviceKey) (err error) {
	return tx.Set(key.MarshalKey(nil), nil)
}
