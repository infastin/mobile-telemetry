package queries

import (
	"github.com/google/uuid"
)

const UserPrefix = "user"

type UserKey struct {
	ID        uuid.UUID
	cachedKey []byte
}

func NewUserKey(id uuid.UUID) *UserKey {
	return &UserKey{
		ID:        id,
		cachedKey: nil,
	}
}

func (u *UserKey) Equal(other *UserKey) bool {
	return u.ID == other.ID
}

func (u *UserKey) MarshalBinary() (data []byte, err error) {
	if u.cachedKey != nil {
		return u.cachedKey, nil
	}

	data = append(data, UserPrefix...)
	data = append(data, ':')
	data = append(data, u.ID[:]...)

	u.cachedKey = data

	return u.cachedKey, nil
}

func (tx *UpdateTx) InsertUser(key *UserKey) (err error) {
	return insertUser(tx, key)
}

func insertUser(tx writeTx, key *UserKey) (err error) {
	keyb, _ := key.MarshalBinary()
	return tx.Set(keyb, nil)
}
