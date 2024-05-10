package impl

import (
	"context"
	"database/sql"
	"mobile-telemetry/pkg/errdefer"
	database "mobile-telemetry/server/service/repo/db"
	"mobile-telemetry/server/service/repo/db/sqliteimpl/sqlc"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type dbRepo struct {
	lg      zerolog.Logger
	db      *sql.DB
	queries *sqlc.Queries
}

type Config struct {
	fx.In

	Logger zerolog.Logger

	Path string `name:"db_path"`
}

func New(lc fx.Lifecycle, cfg Config) (database.Repo, error) {
	db, err := sql.Open("sqlite3", cfg.Path+"?cache=shared&_fk=1&_journal=wal&_sync=normal")
	if err != nil {
		return nil, err
	}
	defer errdefer.Close(&err, db)

	queries, err := sqlc.Prepare(context.Background(), db)
	if err != nil {
		return nil, err
	}

	repo := &dbRepo{
		lg:      cfg.Logger,
		db:      db,
		queries: queries,
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

func (db *dbRepo) atomic(callback database.AtomicCallback) (err error) {
	tx, err := db.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}

		if err != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				err = database.NewRollbackError(err)
			}
		} else {
			err = tx.Commit()
		}
	}()

	newRepo := &dbRepo{
		lg:      db.lg,
		db:      db.db,
		queries: db.queries.WithTx(tx),
	}

	return callback(newRepo)
}

func (db *dbRepo) Update(ctx context.Context, callback database.AtomicCallback) (err error) {
	return db.atomic(callback)
}

func (db *dbRepo) View(ctx context.Context, callback database.AtomicCallback) (err error) {
	return db.atomic(callback)
}

func (db *dbRepo) Batch(ctx context.Context, callback database.AtomicCallback) (err error) {
	return db.atomic(callback)
}

func (db *dbRepo) UserRepo() database.UserRepo {
	return db
}

func (db *dbRepo) DeviceRepo() database.DeviceRepo {
	return db
}

func (db *dbRepo) TelemetryRepo() database.TelemetryRepo {
	return db
}
