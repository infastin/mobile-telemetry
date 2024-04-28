package queries

import (
	"mobile-telemetry/pkg/fastconv"

	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
)

//go:generate msgp -tests=false

const DeviceIndexPrefix = "device_idx"

type DeviceIndexKey struct {
	Manufacturer string `msg:"manufacturer"`
	Model        string `msg:"model"`
	BuildNumber  string `msg:"build_number"`
	cachedKey    []byte
}

func NewDeviceIndexKey(manufacturer, model, buildNumber string) *DeviceIndexKey {
	return &DeviceIndexKey{
		Manufacturer: manufacturer,
		Model:        model,
		BuildNumber:  buildNumber,
		cachedKey:    nil,
	}
}

func (d *DeviceIndexKey) Equal(other *DeviceIndexKey) bool {
	return d.Manufacturer == other.Manufacturer &&
		d.Model == other.Model &&
		d.BuildNumber == other.Model
}

type DeviceIndexIDs []uint64
type DeviceIndexKeys []DeviceIndexKey

func (d *DeviceIndexKey) MarshalBinary() (data []byte, err error) {
	if d.cachedKey != nil {
		return d.cachedKey, nil
	}

	data = append(data, fastconv.Bytes(DeviceIndexPrefix)...)
	data = append(data, ':')

	digest := xxhash.New()
	digest.WriteString(d.Manufacturer)
	digest.WriteString(d.Model)
	digest.WriteString(d.BuildNumber)

	d.cachedKey = digest.Sum(data)

	return d.cachedKey, nil
}

func (tx *UpdateTx) InsertDeviceIndex(idx *DeviceIndexKey) (id uint64, err error) {
	return insertDeviceIndex(tx, tx.queries.deviceSeq, idx)
}

func insertDeviceIndex(inserter Inserter, seq Sequence, idx *DeviceIndexKey) (id uint64, err error) {
	key, _ := idx.MarshalBinary()

	item, err := inserter.Get(key)
	if err != nil && err != badger.ErrKeyNotFound {
		return 0, err
	}

	if err == badger.ErrKeyNotFound {
		return insertDeviceIndexWhenNotFound(inserter, seq, idx, key)
	}

	return insertDeviceIndexWhenExists(inserter, seq, idx, item)
}

func insertDeviceIndexWhenNotFound(inserter Inserter, seq Sequence, idx *DeviceIndexKey, key []byte,
) (id uint64, err error) {
	id, err = seq.Next()
	if err != nil {
		return 0, err
	}

	ids := DeviceIndexIDs{id}
	keys := DeviceIndexKeys{*idx}

	val, _ := ids.MarshalMsg(nil)
	val, _ = keys.MarshalMsg(val)

	err = inserter.Set(key, val)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func insertDeviceIndexWhenExists(inserter Inserter, seq Sequence, idx *DeviceIndexKey, item *badger.Item,
) (id uint64, err error) {
	var (
		ids  DeviceIndexIDs
		keys DeviceIndexKeys
	)

	err = item.Value(func(val []byte) error {
		val, err = ids.UnmarshalMsg(val)
		if err != nil {
			return err
		}
		_, err = keys.UnmarshalMsg(val)
		return err
	})
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(keys); i++ {
		if idx.Equal(&keys[i]) {
			return 0, badger.ErrConflict
		}
	}

	id, err = seq.Next()
	if err != nil {
		return 0, err
	}

	ids = append(ids, id)
	keys = append(keys, *idx)

	val, _ := ids.MarshalMsg(nil)
	val, _ = keys.MarshalMsg(val)

	entry := badger.NewEntry(item.Key(), val).
		WithMeta(byte(Meta(0).SetCollision(true)))

	err = inserter.SetEntry(entry)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tx *ViewTx) GetDeviceIndex(idx *DeviceIndexKey) (id uint64, err error) {
	return getDeviceIndex(tx, idx)
}

func (tx *UpdateTx) GetDeviceIndex(idx *DeviceIndexKey) (id uint64, err error) {
	return getDeviceIndex(tx, idx)
}

func getDeviceIndex(getter Getter, idx *DeviceIndexKey) (id uint64, err error) {
	key, _ := idx.MarshalBinary()

	item, err := getter.Get(key)
	if err != nil {
		return 0, err
	}

	m := Meta(item.UserMeta())
	if m.Collision() {
		return getDeviceIndexOnCollision(idx, item)
	}

	return getDeviceIndexWhenNoCollision(item)
}

func getDeviceIndexOnCollision(idx *DeviceIndexKey, item *badger.Item) (id uint64, err error) {
	var (
		ids  DeviceIndexIDs
		keys DeviceIndexKeys
	)

	err = item.Value(func(val []byte) error {
		val, err = ids.UnmarshalMsg(val)
		if err != nil {
			return err
		}
		_, err = keys.UnmarshalMsg(val)
		return err
	})
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(keys); i++ {
		if idx.Equal(&keys[i]) {
			return ids[i], nil
		}
	}

	return 0, badger.ErrKeyNotFound
}

func getDeviceIndexWhenNoCollision(item *badger.Item) (id uint64, err error) {
	var ids DeviceIndexIDs

	err = item.Value(func(val []byte) error {
		_, err = ids.UnmarshalMsg(val)
		return err
	})

	if err != nil {
		return 0, err
	}

	return ids[0], nil
}
