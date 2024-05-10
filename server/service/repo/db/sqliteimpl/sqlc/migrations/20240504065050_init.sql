-- Create "users" table
CREATE TABLE `users` (`id` text NOT NULL, PRIMARY KEY (`id`));
-- Create "devices" table
CREATE TABLE `devices` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `manufacturer` text NOT NULL, `model` text NOT NULL, `build_number` text NOT NULL, `os` text NOT NULL, `screen_width` integer NOT NULL, `screen_height` integer NOT NULL);
-- Create index "devices_manufacturer_model_build_number" to table: "devices"
CREATE UNIQUE INDEX `devices_manufacturer_model_build_number` ON `devices` (`manufacturer`, `model`, `build_number`);
-- Create "user_devices" table
CREATE TABLE `user_devices` (`user_id` text NOT NULL, `device_id` integer NOT NULL, PRIMARY KEY (`user_id`, `device_id`), CONSTRAINT `0` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "telemetries" table
CREATE TABLE `telemetries` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `user_id` text NOT NULL, `device_id` integer NOT NULL, `os_version` text NOT NULL, `app_version` text NOT NULL, `action` text NOT NULL, `data` blob NOT NULL, `timestamp` datetime NOT NULL, CONSTRAINT `0` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
