package impl

import (
	"context"
	"mobile-telemetry/pkg/errdefer"
	database "mobile-telemetry/server/service/repo/db"
	"mobile-telemetry/server/service/repo/db/badgerimpl/schema"
	"time"

	"github.com/dgraph-io/badger/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type dbRepo struct {
	lg *zap.Logger
	db *badger.DB

	deviceSeq *badger.Sequence
}

type Config struct {
	fx.In

	Logger *zap.Logger

	Directory string `name:"db_directory"`
}

func New(lc fx.Lifecycle, cfg Config) (database.Repo, error) {
	db, err := badger.Open(badger.DefaultOptions(cfg.Directory))
	if err != nil {
		return nil, err
	}
	defer errdefer.Close(&err, db)

	deviceSeq, err := db.GetSequence([]byte(schema.DevicePrefix), 256)
	if err != nil {
		return nil, err
	}

	repo := &dbRepo{
		lg:        cfg.Logger,
		db:        db,
		deviceSeq: deviceSeq,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				ticker := time.NewTicker(5 * time.Minute)
				defer ticker.Stop()
				for {
					select {
					case <-ctx.Done():
						return
					case <-ticker.C:
						_ = repo.db.RunValueLogGC(0.5)
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return repo.Close()
		},
	})

	return repo, nil
}

func (db *dbRepo) Close() error {
	_ = db.deviceSeq.Release()
	return db.db.Close()
}

func (db *dbRepo) Atomic(ctx context.Context, callback database.AtomicCallback) (err error) {
	return callback(db)
}

func (db *dbRepo) User() database.UserRepo {
	return db
}

func (db *dbRepo) Device() database.DeviceRepo {
	return db
}

func (db *dbRepo) Telemetry() database.TelemetryRepo {
	return db
}
