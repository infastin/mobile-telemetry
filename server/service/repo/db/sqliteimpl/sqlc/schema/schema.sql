CREATE TABLE users (
  id text NOT NULL PRIMARY KEY
);

CREATE TABLE devices (
  id            integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  manufacturer  text    NOT NULL,
  model         text    NOT NULL,
  build_number  text    NOT NULL,
  os            text    NOT NULL,
  screen_width  integer NOT NULL,
  screen_height integer NOT NULL,
  UNIQUE (manufacturer, model, build_number)
);

CREATE TABLE user_devices (
  user_id   text    NOT NULL REFERENCES users(id),
  device_id integer NOT NULL REFERENCES devices(id),
  PRIMARY KEY (user_id, device_id)
);

CREATE TABLE telemetries (
  id          integer  NOT NULL PRIMARY KEY AUTOINCREMENT,
  user_id     text     NOT NULL REFERENCES users(id),
  device_id   integer  NOT NULL REFERENCES devices(id),
  os_version  text     NOT NULL,
  app_version text     NOT NULL,
  action      text     NOT NULL,
  data        blob     NOT NULL,
  timestamp   datetime NOT NULL
);
