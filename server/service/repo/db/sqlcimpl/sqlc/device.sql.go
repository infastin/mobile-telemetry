// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: device.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

const findDevice = `-- name: FindDevice :one
SELECT id, manufacturer, model, build_number, os, screen_width, screen_height FROM devices WHERE manufacturer = $1 AND model = $2 AND build_number = $3
`

type FindDeviceParams struct {
	Manufacturer string
	Model        string
	BuildNumber  string
}

func (q *Queries) FindDevice(ctx context.Context, arg FindDeviceParams) (Device, error) {
	row := q.db.QueryRow(ctx, findDevice, arg.Manufacturer, arg.Model, arg.BuildNumber)
	var i Device
	err := row.Scan(
		&i.ID,
		&i.Manufacturer,
		&i.Model,
		&i.BuildNumber,
		&i.Os,
		&i.ScreenWidth,
		&i.ScreenHeight,
	)
	return i, err
}

const upsertDevice = `-- name: UpsertDevice :execresult
INSERT INTO devices (manufacturer, model, build_number, os, screen_width, screen_height) VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT DO NOTHING
`

type UpsertDeviceParams struct {
	Manufacturer string
	Model        string
	BuildNumber  string
	Os           string
	ScreenWidth  int32
	ScreenHeight int32
}

func (q *Queries) UpsertDevice(ctx context.Context, arg UpsertDeviceParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, upsertDevice,
		arg.Manufacturer,
		arg.Model,
		arg.BuildNumber,
		arg.Os,
		arg.ScreenWidth,
		arg.ScreenHeight,
	)
}
