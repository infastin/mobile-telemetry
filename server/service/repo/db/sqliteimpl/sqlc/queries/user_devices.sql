-- name: InsertUserDeviceIfNotExists :execresult
INSERT INTO user_devices (user_id, device_id) VALUES (?, ?)
ON CONFLICT DO NOTHING;
