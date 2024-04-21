package app

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/fx"
	"go.uber.org/zap/zapcore"
)

type (
	Config struct {
		fx.Out

		Application ApplicationConfig `env-prefix:"APP_" yaml:"app"`
		Logger      LoggerConfig      `env-prefix:"LOG_" yaml:"log"`
		Database    DatabaseConfig    `env-prefix:"DB_" yaml:"db"`
		HTTP        HTTPConfig        `env-prefix:"HTTP_" yaml:"http"`
	}

	ApplicationConfig struct {
		fx.Out

		Mode Mode `env-default:"debug" yaml:"mode" name:"app_mode"`
	}

	LoggerConfig struct {
		fx.Out

		Level struct {
			fx.Out

			Debug   zapcore.Level `env-default:"debug" yaml:"debug" name:"log_level_debug"`
			Release zapcore.Level `env-default:"info" yaml:"release" name:"log_level_release"`
		} `env-prefix:"LEVEL_" yaml:"level"`
	}

	DatabaseConfig struct {
		fx.Out

		Host     string `env-required:"" env:"HOST" yaml:"host" name:"db_host"`
		Port     int    `env-required:"" env:"PORT" yaml:"port" name:"db_port"`
		User     string `env-required:"" env:"USER" yaml:"user" name:"db_user"`
		Password string `env-required:"" env:"PASSWORD" yaml:"password" name:"db_password"`
		Name     string `env-required:"" env:"NAME" yaml:"name" name:"db_name"`
		SSLMode  string `env-required:"" env:"SSLMODE" yaml:"sslmode" name:"db_sslmode"`
	}

	HTTPConfig struct {
		fx.Out

		Port int `env-required:"" env:"PORT" yaml:"port" name:"http_port"`
	}
)

func NewConfig(configPath string) (cfg Config, err error) {
	if configPath != "" {
		err = cleanenv.ReadConfig(configPath, &cfg)
		if err != nil {
			return Config{}, fmt.Errorf("could not read config: %w", err)
		}
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("could not read envvars: %w", err)
	}

	SetMode(cfg.Application.Mode)

	return cfg, nil
}
