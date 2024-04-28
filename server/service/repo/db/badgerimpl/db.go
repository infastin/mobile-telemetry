package impl

import (
	"context"
	"mobile-telemetry/pkg/errdefer"
	"mobile-telemetry/pkg/fastconv"
	database "mobile-telemetry/server/service/repo/db"
	"mobile-telemetry/server/service/repo/db/badgerimpl/queries"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type dbRepo struct {
	lg           zerolog.Logger
	db           *badger.DB
	deviceSeq    *badger.Sequence
	telemetrySeq *badger.Sequence
	queries      *queries.Queries
}

type Config struct {
	fx.In

	Logger       zerolog.Logger
	BadgerLogger badger.Logger

	Directory string `name:"db_directory"`
}

func New(lc fx.Lifecycle, cfg Config) (database.Repo, error) {
	db, err := badger.Open(badger.DefaultOptions(cfg.Directory).
		WithLogger(cfg.BadgerLogger))
	if err != nil {
		return nil, err
	}
	defer errdefer.Close(&err, db)

	deviceSeq, err := db.GetSequence(fastconv.Bytes(queries.DevicePrefix), 256)
	if err != nil {
		return nil, err
	}
	defer errdefer.Release(&err, deviceSeq)

	telemetrySeq, err := db.GetSequence(fastconv.Bytes(queries.TelemetryPrefix), 1024)
	if err != nil {
		return nil, err
	}

	repo := &dbRepo{
		lg:           cfg.Logger,
		db:           db,
		deviceSeq:    deviceSeq,
		telemetrySeq: telemetrySeq,
		queries:      queries.New(db, deviceSeq, telemetrySeq),
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
	_ = db.telemetrySeq.Release()
	return db.db.Close()
}
