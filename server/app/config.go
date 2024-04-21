package app

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/fx"
)

type (
	Config struct {
		fx.Out

		Database DatabaseConfig `env-prefix:"DB_" yaml:"db"`
		HTTP     HTTPConfig     `env-prefix:"HTTP_" yaml:"http"`
	}

	DatabaseConfig struct {
		fx.Out

		Host     string `env-required:"" env:"HOST" yaml:"host" name:"db_host"`
		Port     int    `env-required:"" env:"PORT" yaml:"port" name:"db_port"`
		User     string `env-required:"" env:"USER" yaml:"user" name:"db_user"`
		Password string `env-required:"" env:"PASSWORD" yaml:"password" name:"db_password"`
		Name     string `env-required:"" env:"NAME" yaml:"name" name:"db_name"`
		SSLMode  string `env-required:"" env:"SSL_MODE" yaml:"sslmode" name:"db_sslmode"`
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

	return cfg, nil
}
