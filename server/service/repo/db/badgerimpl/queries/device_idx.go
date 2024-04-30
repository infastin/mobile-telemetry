package queries

import (
	"encoding/binary"
	"mobile-telemetry/pkg/fastconv"
	"slices"

	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
)

//go:generate msgp -tests=false
//msgp:ignore DeviceIndexKey

const DeviceIndexPrefix = "device_idx"

type DeviceIndexKey struct {
	Hash      uint64
	ID        uint64
	Collision bool
}

func NewDeviceIndexKey(manufacturer, model, buildNumber string) *DeviceIndexKey {
	digest := xxhash.New()
	digest.WriteString(manufacturer)
	digest.WriteString(model)
	digest.WriteString(buildNumber)

	return &DeviceIndexKey{
		Hash:      digest.Sum64(),
		ID:        0,
		Collision: false,
	}
}

func (d *DeviceIndexKey) MarshalKey(b []byte) []byte {
	b = slices.Grow(b, len(DeviceIndexPrefix)+1+8+8)
	b = append(b, fastconv.Bytes(DeviceIndexPrefix)...)
	b = append(b, ':')
	b = binary.BigEndian.AppendUint64(b, d.Hash)
	b = binary.BigEndian.AppendUint64(b, d.ID)
	return b
}

func (d *DeviceIndexKey) UnmarshalKey(b []byte) error {
	if len(b) != len(DeviceIndexPrefix)+1+8+8 {
		return NewInvalidKeySizeError(len(DeviceIndexPrefix)+1+8+8, len(b))
	}

	prefix := fastconv.String(b[:len(DeviceIndexPrefix)])
	if prefix != DeviceIndexPrefix {
		return NewInvalidKeyPrefix(DeviceIndexPrefix, prefix)
	}

	b = b[len(DeviceIndexPrefix)+1:]
	d.Hash = binary.BigEndian.Uint64(b)
	b = b[8:]
	d.ID = binary.BigEndian.Uint64(b)

	return nil
}

func (d *DeviceIndexKey) MarshalPrefix(b []byte) []byte {
	b = slices.Grow(b, len(DeviceIndexPrefix)+1+8)
	b = append(b, fastconv.Bytes(DeviceIndexPrefix)...)
	b = append(b, ':')
	b = binary.BigEndian.AppendUint64(b, d.Hash)
	return b
}

type DeviceIndexValue struct {
	Manufacturer string `msg:"manufacturer"`
	Model        string `msg:"model"`
	BuildNumber  string `msg:"build_number"`
}

func (tx *UpdateTx) InsertDeviceIndex(key *DeviceIndexKey, val *DeviceIndexValue) (id uint64, err error) {
	return insertDeviceIndex(tx, tx.queries.deviceSeq, key, val)
}

func insertDeviceIndex(tx updateTx, seq *badger.Sequence, key *DeviceIndexKey, val *DeviceIndexValue,
) (id uint64, err error) {
	key.ID, err = seq.Next()
	if err != nil {
		return 0, err
	}

	keyb := key.MarshalKey(nil)
	valb, _ := val.MarshalMsg(nil)

	err = tx.SetEntry(badger.NewEntry(keyb, valb).
		WithMeta(byte(Meta(0).SetCollision(key.Collision))))
	if err != nil {
		return 0, err
	}

	return key.ID, nil
}

func (tx *ViewTx) FindDeviceIndex(key *DeviceIndexKey, val *DeviceIndexValue) (id uint64, err error) {
	return findDeviceIndex(tx, key, val)
}

func (tx *UpdateTx) FindDeviceIndex(key *DeviceIndexKey, val *DeviceIndexValue) (id uint64, err error) {
	return findDeviceIndex(tx, key, val)
}

func findDeviceIndex(tx viewTx, key *DeviceIndexKey, val *DeviceIndexValue) (id uint64, err error) {
	opts := badger.DefaultIteratorOptions
	opts.PrefetchValues = false
	opts.Prefix = key.MarshalPrefix(nil)

	it := tx.NewIterator(opts)
	defer it.Close()

	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()

		meta := Meta(item.UserMeta())
		if meta.Collision() {
			var itVal DeviceIndexValue
			if err := item.Value(func(valb []byte) error {
				_, err = itVal.UnmarshalMsg(valb)
				return err
			}); err != nil {
				return 0, err
			}

			if itVal != *val {
				continue
			}
		}

		err = key.UnmarshalKey(item.Key())
		if err != nil {
			return 0, err
		}

		key.Collision = meta.Collision()
		return key.ID, nil
	}

	return 0, badger.ErrKeyNotFound
}
