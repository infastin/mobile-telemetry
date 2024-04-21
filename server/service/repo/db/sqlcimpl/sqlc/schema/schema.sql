CREATE TABLE users (
  id uuid NOT NULL PRIMARY KEY
);

CREATE TABLE devices (
  id            serial8 NOT NULL PRIMARY KEY,
  manufacturer  text    NOT NULL,
  model         text    NOT NULL,
  build_number  text    NOT NULL,
  os            text    NOT NULL,
  screen_width  int4    NOT NULL,
  screen_height int4    NOT NULL,
  UNIQUE(manufacturer, model, build_number)
);

CREATE TABLE user_devices (
  user_id   uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  device_id int8 NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
  UNIQUE(user_id, device_id)
);

CREATE TABLE telemetries (
  id          serial8     NOT NULL PRIMARY KEY,
  user_id     uuid        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  device_id   int8        NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
  app_version text        NOT NULL,
  os_version  text        NOT NULL,
  action      text        NOT NULL,
  data        jsonb       NOT NULL,
  timestamp   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX telemetries_user_id_device_id_idx ON telemetries (user_id, device_id);
CREATE INDEX telemetries_user_id_device_id_app_version_idx ON telemetries (user_id, device_id, app_version);
CREATE INDEX telemetries_user_id_device_id_os_version_idx ON telemetries (user_id, device_id, os_version);
