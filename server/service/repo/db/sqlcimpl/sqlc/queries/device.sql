-- name: UpserDevice :execresult
INSERT INTO devices (manufacturer, model, build_number, os, screen_width, screen_height) VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT DO NOTHING;

-- name: FindDevice :one
SELECT * FROM devices WHERE manufacturer = $1 AND model = $2 AND build_number = $3;
