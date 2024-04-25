package schema

import (
	"bytes"
	"encoding/binary"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
)

const UserPrefix string = "user"

type User struct {
	ID uuid.UUID
}

func UserKey(id uuid.UUID) []byte {
	var b bytes.Buffer

	b.Grow(len(UserPrefix) + 1 + len(id))
	b.WriteString(UserPrefix)
	b.WriteByte(':')
	_ = binary.Write(&b, binary.BigEndian, id)

	return b.Bytes()
}

func MarshalUserData(data *UserData) ([]byte, error) {
	return data.MarshalMsg(nil)
}

func UserEntry(user *User) (*badger.Entry, error) {
	key := UserKey(user.ID)

	val, err := MarshalUserData(&UserData{})
	if err != nil {
		return nil, err
	}

	return badger.NewEntry(key, val), nil
}
