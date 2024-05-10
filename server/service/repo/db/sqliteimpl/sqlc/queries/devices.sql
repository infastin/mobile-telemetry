-- name: InsertDevice :execresult
INSERT INTO devices (manufacturer, model, build_number, os, screen_width, screen_height) VALUES (?, ?, ?, ?, ?, ?);

-- name: UpsertDevice :execresult
INSERT INTO devices (manufacturer, model, build_number, os, screen_width, screen_height) VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT DO NOTHING;

-- name: FindDevice :one
SELECT * FROM devices WHERE manufacturer = ? AND model = ? AND build_number = ?;
