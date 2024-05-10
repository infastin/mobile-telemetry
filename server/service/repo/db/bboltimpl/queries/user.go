package queries

import (
	"slices"

	"github.com/google/uuid"
)

var UserBucketName = []byte("user")

type UserKey struct {
	ID uuid.UUID
}

func NewUserKey(id uuid.UUID) *UserKey {
	return &UserKey{
		ID: id,
	}
}

func (u *UserKey) MarshalKey(b []byte) []byte {
	b = slices.Grow(b, 16)
	b = append(b, u.ID[:]...)
	return b
}

func (queries *Queries) InsertUser(key *UserKey) (err error) {
	b := queries.tx.Bucket(UserBucketName)

	keyb := key.MarshalKey(nil)
	valb := Meta(0).Append(nil)

	return b.Put(keyb, valb)
}
