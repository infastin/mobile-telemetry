package queries

import (
	"encoding/binary"

	"github.com/google/uuid"
)

const UserDevicePrefix = "user_device"

type UserDeviceKey struct {
	UserID    uuid.UUID
	DeviceID  uint64
	cachedKey []byte
}

func NewUserDeviceKey(userID uuid.UUID, deviceID uint64) *UserDeviceKey {
	return &UserDeviceKey{
		UserID:    userID,
		DeviceID:  deviceID,
		cachedKey: nil,
	}
}

func (ud *UserDeviceKey) Equal(other *UserDeviceKey) bool {
	return ud.UserID == other.UserID && ud.DeviceID == other.DeviceID
}

func (ud *UserDeviceKey) MarshalBinary() (data []byte, err error) {
	if ud.cachedKey != nil {
		return ud.cachedKey, nil
	}

	data = make([]byte, 0, len(UserDevicePrefix)+1+16+1+8)
	data = append(data, UserDevicePrefix...)
	data = append(data, ':')
	data = append(data, ud.UserID[:]...)
	data = append(data, ':')
	data = binary.BigEndian.AppendUint64(data, ud.DeviceID)

	ud.cachedKey = data

	return ud.cachedKey, nil
}

func (tx *UpdateTx) InsertUserDevice(key *UserDeviceKey) (err error) {
	return insertUserDevice(tx, key)
}

func insertUserDevice(tx writeTx, key *UserDeviceKey) (err error) {
	keyb, _ := key.MarshalBinary()
	return tx.Set(keyb, nil)
}
