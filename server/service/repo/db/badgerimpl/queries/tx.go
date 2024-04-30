package queries

import (
	"github.com/dgraph-io/badger/v4"
)

type writeTx interface {
	Set(key, val []byte) error
	SetEntry(entry *badger.Entry) error
}

type viewTx interface {
	Get(key []byte) (item *badger.Item, err error)
}

type updateTx interface {
	writeTx
	viewTx
}

type UpdateTx struct {
	queries *Queries
	tx      *badger.Txn
}

func (tx *UpdateTx) Set(key, val []byte) (err error) {
	return tx.SetEntry(badger.NewEntry(key, val))
}

func (tx *UpdateTx) SetEntry(entry *badger.Entry) (err error) {
	err = tx.tx.SetEntry(entry)
	if err != nil && err != badger.ErrTxnTooBig {
		return err
	}

	if err != badger.ErrTxnTooBig {
		return nil
	}

	_ = tx.tx.Commit()
	tx.tx = tx.queries.db.NewTransaction(true)

	return tx.tx.SetEntry(entry)
}

func (tx *UpdateTx) Get(key []byte) (item *badger.Item, err error) {
	return tx.tx.Get(key)
}

func (tx *UpdateTx) Commit() (err error) {
	return tx.tx.Commit()
}

func (tx *UpdateTx) Discard() {
	tx.tx.Discard()
}

type BatchWriteTx struct {
	queries *Queries
	tx      *badger.WriteBatch
}

func (tx *BatchWriteTx) Set(key, val []byte) (err error) {
	return tx.tx.Set(key, val)
}

func (tx *BatchWriteTx) SetEntry(entry *badger.Entry) (err error) {
	return tx.tx.SetEntry(entry)
}

func (tx *BatchWriteTx) Commit() (err error) {
	return tx.tx.Flush()
}

func (tx *BatchWriteTx) Discard() {
	tx.tx.Cancel()
}

type ViewTx struct {
	tx *badger.Txn
}

func (tx *ViewTx) Get(key []byte) (item *badger.Item, err error) {
	return tx.tx.Get(key)
}

func (tx *ViewTx) Commit() (err error) {
	return tx.tx.Commit()
}

func (tx *ViewTx) Discard() {
	tx.tx.Discard()
}
