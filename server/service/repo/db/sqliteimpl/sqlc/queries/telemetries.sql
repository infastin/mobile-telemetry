-- name: InsertTelemetry :execresult
INSERT INTO telemetries (user_id, device_id, os_version, app_version, action, data, timestamp) VALUES (?, ?, ?, ?, ?, ?, ?);
