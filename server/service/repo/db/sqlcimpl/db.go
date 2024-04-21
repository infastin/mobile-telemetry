package impl

import (
	"context"
	database "mobile-telemetry/server/service/repo/db"
	"mobile-telemetry/server/service/repo/db/sqlcimpl/sqlc"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type dbRepo struct {
	lg      *zap.Logger
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

type Config struct {
	fx.In

	Logger *zap.Logger

	Host     string `name:"db_host"`
	Port     int    `name:"db_port"`
	User     string `name:"db_user"`
	Password string `name:"db_password"`
	Name     string `name:"db_name"`
	SSLMode  string `name:"db_sslmode"`
}

func New(lc fx.Lifecycle, cfg Config) (database.Repo, error) {
	params := map[string]string{
		"host":     cfg.Host,
		"port":     strconv.Itoa(cfg.Port),
		"user":     cfg.User,
		"password": cfg.Password,
		"dbname":   cfg.Name,
		"sslmode":  cfg.SSLMode,
	}

	var dsn strings.Builder
	for k, v := range params {
		dsn.WriteString(k)
		dsn.WriteByte('=')
		dsn.WriteString(v)
		dsn.WriteByte(' ')
	}

	pool, err := pgxpool.New(context.Background(), dsn.String())
	if err != nil {
		return nil, err
	}

	repo := &dbRepo{
		lg:      cfg.Logger,
		pool:    pool,
		queries: sqlc.New(pool),
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			repo.Close()
			return nil
		},
	})

	return repo, nil
}

func (db *dbRepo) Close() {
	db.pool.Close()
}

func (db *dbRepo) Atomic(ctx context.Context, callback database.AtomicCallback) (err error) {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		}

		if err != nil {
			rbErr := tx.Rollback(ctx)
			if rbErr != nil {
				err = database.NewRollbackError(err)
			}
		} else {
			err = tx.Commit(ctx)
		}
	}()

	newRepo := &dbRepo{
		lg:      db.lg,
		pool:    db.pool,
		queries: db.queries.WithTx(tx),
	}

	return callback(newRepo)
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
