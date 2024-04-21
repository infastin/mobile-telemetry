-- name: InsertTelemetriesBulk :copyfrom
INSERT INTO telemetries (user_id, device_id, os_version, app_version, action, data, timestamp)
  VALUES ($1, $2, $3, $4, $5, $6, $7);
