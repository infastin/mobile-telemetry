// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: devices.sql

package sqlc

import (
	"context"
	"database/sql"
)

const findDevice = `-- name: FindDevice :one
SELECT id, manufacturer, model, build_number, os, screen_width, screen_height FROM devices WHERE manufacturer = ? AND model = ? AND build_number = ?
`

type FindDeviceParams struct {
	Manufacturer string
	Model        string
	BuildNumber  string
}

func (q *Queries) FindDevice(ctx context.Context, arg FindDeviceParams) (Device, error) {
	row := q.queryRow(ctx, q.findDeviceStmt, findDevice, arg.Manufacturer, arg.Model, arg.BuildNumber)
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

const insertDevice = `-- name: InsertDevice :execresult
INSERT INTO devices (manufacturer, model, build_number, os, screen_width, screen_height) VALUES (?, ?, ?, ?, ?, ?)
`

type InsertDeviceParams struct {
	Manufacturer string
	Model        string
	BuildNumber  string
	Os           string
	ScreenWidth  int64
	ScreenHeight int64
}

func (q *Queries) InsertDevice(ctx context.Context, arg InsertDeviceParams) (sql.Result, error) {
	return q.exec(ctx, q.insertDeviceStmt, insertDevice,
		arg.Manufacturer,
		arg.Model,
		arg.BuildNumber,
		arg.Os,
		arg.ScreenWidth,
		arg.ScreenHeight,
	)
}

const upsertDevice = `-- name: UpsertDevice :execresult
INSERT INTO devices (manufacturer, model, build_number, os, screen_width, screen_height) VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT DO NOTHING
`

type UpsertDeviceParams struct {
	Manufacturer string
	Model        string
	BuildNumber  string
	Os           string
	ScreenWidth  int64
	ScreenHeight int64
}

func (q *Queries) UpsertDevice(ctx context.Context, arg UpsertDeviceParams) (sql.Result, error) {
	return q.exec(ctx, q.upsertDeviceStmt, upsertDevice,
		arg.Manufacturer,
		arg.Model,
		arg.BuildNumber,
		arg.Os,
		arg.ScreenWidth,
		arg.ScreenHeight,
	)
}
