package queries

import (
	"bytes"
	"encoding/binary"
	"slices"

	"github.com/cespare/xxhash/v2"
)

//go:generate msgp -tests=false
//msgp:ignore DeviceIndexKey

var DeviceIndexBucketName = []byte("device_index")

type DeviceIndexKey struct {
	Hash      uint64
	ID        uint64
	Collision bool
}

func NewDeviceIndexKey(manufacturer, model, buildNumber string) *DeviceIndexKey {
	digest := xxhash.New()
	_, _ = digest.WriteString(manufacturer)
	_, _ = digest.WriteString(model)
	_, _ = digest.WriteString(buildNumber)

	return &DeviceIndexKey{
		Hash:      digest.Sum64(),
		ID:        0,
		Collision: false,
	}
}

func (d *DeviceIndexKey) MarshalKey(b []byte) []byte {
	b = slices.Grow(b, 8+8)
	b = binary.BigEndian.AppendUint64(b, d.Hash)
	b = binary.BigEndian.AppendUint64(b, d.ID)
	return b
}

func (d *DeviceIndexKey) UnmarshalKey(b []byte) error {
	if len(b) != 8+8 {
		return NewInvalidKeySizeError(8+8, len(b))
	}

	d.Hash = binary.BigEndian.Uint64(b)
	d.ID = binary.BigEndian.Uint64(b[8:])

	return nil
}

func (d *DeviceIndexKey) MarshalPrefix(b []byte) []byte {
	b = slices.Grow(b, 8)
	b = binary.BigEndian.AppendUint64(b, d.Hash)
	return b
}

type DeviceIndexValue struct {
	Manufacturer string `msg:"manufacturer"`
	Model        string `msg:"model"`
	BuildNumber  string `msg:"build_number"`
}

func (queries *Queries) InsertDeviceIndex(key *DeviceIndexKey, val *DeviceIndexValue,
) (id uint64, err error) {
	b := queries.tx.Bucket(DeviceIndexBucketName)

	key.ID, err = b.NextSequence()
	if err != nil {
		return 0, err
	}

	keyb := key.MarshalKey(nil)

	valb := Meta(0).SetCollision(key.Collision).Append(nil)
	valb, _ = val.MarshalMsg(valb)

	err = b.Put(keyb, valb)
	if err != nil {
		return 0, err
	}

	return key.ID, nil
}

func (queries *Queries) FindDeviceIndex(key *DeviceIndexKey, val *DeviceIndexValue) (id uint64, err error) {
	b := queries.tx.Bucket(DeviceIndexBucketName)

	c := b.Cursor()
	prefix := key.MarshalPrefix(nil)

	for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
		meta := Meta(v[0])
		if meta.Collision() {
			var cVal DeviceIndexValue

			_, err = cVal.UnmarshalMsg(v[1:])
			if err != nil {
				return 0, err
			}

			if cVal != *val {
				continue
			}
		}

		err = key.UnmarshalKey(k)
		if err != nil {
			return 0, err
		}

		key.Collision = meta.Collision()
		return key.ID, nil
	}

	return 0, ErrKeyNotFound
}
