package impl

import (
	"context"
	"database/sql"
	database "mobile-telemetry/server/service/repo/db"
	"mobile-telemetry/server/service/repo/db/entimpl/ent"
	"strconv"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type dbRepo struct {
	lg     *zap.Logger
	client *ent.Client
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

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	repo := &dbRepo{
		lg:     cfg.Logger,
		client: ent.NewClient(ent.Driver(drv)),
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return repo.Close()
		},
	})

	return repo, nil
}

func (db *dbRepo) Close() (err error) {
	return db.client.Close()
}

func (db *dbRepo) Atomic(ctx context.Context, callback database.AtomicCallback) error {
	tx, err := db.client.BeginTx(ctx, &sql.TxOptions{})
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
				err = database.NewRollbackError(rbErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	newRepo := &dbRepo{
		lg:     db.lg,
		client: tx.Client(),
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
