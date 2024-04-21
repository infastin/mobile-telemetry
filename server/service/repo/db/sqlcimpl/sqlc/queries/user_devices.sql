-- name: UpsertUserDevice :execresult
INSERT INTO user_devices (user_id, device_id) VALUES ($1, $2)
ON CONFLICT DO NOTHING;
