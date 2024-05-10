package impl

import (
	"context"
	"mobile-telemetry/pkg/errdefer"
	database "mobile-telemetry/server/service/repo/db"
	"mobile-telemetry/server/service/repo/db/bboltimpl/queries"

	"github.com/rs/zerolog"
	"go.etcd.io/bbolt"
	"go.uber.org/fx"
)

type dbRepo struct {
	lg      zerolog.Logger
	db      *bbolt.DB
	queries *queries.Queries
}

type Config struct {
	fx.In

	Logger zerolog.Logger

	Path string `name:"db_path"`
}

func New(lc fx.Lifecycle, cfg Config) (database.Repo, error) {
	db, err := bbolt.Open(cfg.Path, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer errdefer.Close(&err, db)

	err = queries.Prepare(db)
	if err != nil {
		return nil, err
	}

	repo := &dbRepo{
		lg:      cfg.Logger,
		db:      db,
		queries: nil,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return repo.Close()
		},
	})

	return repo, nil
}

func (db *dbRepo) Close() error {
	return db.db.Close()
}

func (db *dbRepo) atomic(callback database.AtomicCallback, tx *bbolt.Tx) (err error) {
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	newRepo := &dbRepo{
		lg:      db.lg,
		db:      db.db,
		queries: queries.New(tx),
	}

	return callback(newRepo)
}

func (db *dbRepo) Update(ctx context.Context, callback database.AtomicCallback) (err error) {
	return db.db.Update(func(tx *bbolt.Tx) error {
		return db.atomic(callback, tx)
	})
}

func (db *dbRepo) View(ctx context.Context, callback database.AtomicCallback) (err error) {
	return db.db.View(func(tx *bbolt.Tx) error {
		return db.atomic(callback, tx)
	})
}

func (db *dbRepo) Batch(ctx context.Context, callback database.AtomicCallback) (err error) {
	return db.db.Batch(func(tx *bbolt.Tx) error {
		return db.atomic(callback, tx)
	})
}

func (db *dbRepo) UserRepo() database.UserRepo {
	if db.queries == nil {
		panic(ErrTxNotStarted)
	}
	return db
}

func (db *dbRepo) DeviceRepo() database.DeviceRepo {
	if db.queries == nil {
		panic(ErrTxNotStarted)
	}
	return db
}

func (db *dbRepo) TelemetryRepo() database.TelemetryRepo {
	if db.queries == nil {
		panic(ErrTxNotStarted)
	}
	return db
}
