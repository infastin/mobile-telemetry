package queries

import (
	"slices"

	"github.com/google/uuid"
)

const UserPrefix = "user"

type UserKey struct {
	ID uuid.UUID
}

func NewUserKey(id uuid.UUID) *UserKey {
	return &UserKey{
		ID: id,
	}
}

func (u *UserKey) MarshalKey(b []byte) []byte {
	b = slices.Grow(b, len(UserPrefix)+1+16)
	b = append(b, UserPrefix...)
	b = append(b, ':')
	b = append(b, u.ID[:]...)
	return b
}

func (tx *UpdateTx) InsertUser(key *UserKey) (err error) {
	return insertUser(tx, key)
}

func insertUser(tx writeTx, key *UserKey) (err error) {
	return tx.Set(key.MarshalKey(nil), nil)
}
