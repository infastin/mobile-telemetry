package queries

import (
	"github.com/dgraph-io/badger/v4"
)

type Queries struct {
	db           *badger.DB
	deviceSeq    *badger.Sequence
	telemetrySeq *badger.Sequence
}

func New(db *badger.DB, deviceSeq, telemetrySeq *badger.Sequence) *Queries {
	return &Queries{
		db:           db,
		deviceSeq:    deviceSeq,
		telemetrySeq: telemetrySeq,
	}
}

func (q *Queries) Update() *UpdateTx {
	return &UpdateTx{
		queries: q,
		tx:      q.db.NewTransaction(true),
	}
}

func (q *Queries) BatchWrite() *BatchWriteTx {
	return &BatchWriteTx{
		queries: q,
		tx:      q.db.NewWriteBatch(),
	}
}

func (q *Queries) View() *ViewTx {
	return &ViewTx{
		tx: q.db.NewTransaction(false),
	}
}
