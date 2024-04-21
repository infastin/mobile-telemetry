-- Create "devices" table
CREATE TABLE "devices" ("id" bigserial NOT NULL, "manufacturer" text NOT NULL, "model" text NOT NULL, "build_number" text NOT NULL, "os" text NOT NULL, "screen_width" integer NOT NULL, "screen_height" integer NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "devices_manufacturer_model_build_number_key" UNIQUE ("manufacturer", "model", "build_number"));
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, PRIMARY KEY ("id"));
-- Create "telemetries" table
CREATE TABLE "telemetries" ("id" bigserial NOT NULL, "user_id" uuid NOT NULL, "device_id" bigint NOT NULL, "app_version" text NOT NULL, "os_version" text NOT NULL, "action" text NOT NULL, "data" jsonb NOT NULL, "timestamp" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "telemetries_device_id_fkey" FOREIGN KEY ("device_id") REFERENCES "devices" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "telemetries_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "telemetries_user_id_device_id_app_version_idx" to table: "telemetries"
CREATE INDEX "telemetries_user_id_device_id_app_version_idx" ON "telemetries" ("user_id", "device_id", "app_version");
-- Create index "telemetries_user_id_device_id_idx" to table: "telemetries"
CREATE INDEX "telemetries_user_id_device_id_idx" ON "telemetries" ("user_id", "device_id");
-- Create index "telemetries_user_id_device_id_os_version_idx" to table: "telemetries"
CREATE INDEX "telemetries_user_id_device_id_os_version_idx" ON "telemetries" ("user_id", "device_id", "os_version");
-- Create "user_devices" table
CREATE TABLE "user_devices" ("user_id" uuid NOT NULL, "device_id" bigint NOT NULL, CONSTRAINT "user_devices_user_id_device_id_key" UNIQUE ("user_id", "device_id"), CONSTRAINT "user_devices_device_id_fkey" FOREIGN KEY ("device_id") REFERENCES "devices" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "user_devices_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
