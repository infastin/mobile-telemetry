package app

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/infastin/go-validation"
	isint "github.com/infastin/go-validation/is/int"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type Config struct {
	fx.Out

	Logger   LoggerConfig   `env-prefix:"LOGGER_" yaml:"logger"`
	Database DatabaseConfig `env-prefix:"DATABASE_" yaml:"database"`
	HTTP     HTTPConfig     `env-prefix:"HTTP_" yaml:"http"`
}

func (cfg Config) Validate() error {
	return validation.All(
		validation.Ptr(&cfg.Logger, "logger").With(validation.Custom),
		validation.Ptr(&cfg.Database, "database").With(validation.Custom),
		validation.Ptr(&cfg.HTTP, "http").With(validation.Custom),
	)
}

type LoggerConfig struct {
	fx.Out

	Level      zerolog.Level `env:"LEVEL" yaml:"level" name:"logger_level"`
	Directory  string        `env:"DIRECTORY" yaml:"directory" name:"logger_directory"`
	MaxSize    int           `env:"MAX_SIZE" yaml:"max_size" name:"logger_max_size"`
	MaxAge     int           `env:"MAX_AGE" yaml:"max_age" name:"logger_max_age"`
	MaxBackups int           `env:"MAX_BACKUPS" yaml:"max_backups" name:"logger_max_backups"`
}

func (cfg LoggerConfig) Validate() error {
	return validation.All(
		validation.Comparable(cfg.Level, "level").In(
			zerolog.TraceLevel, zerolog.DebugLevel,
			zerolog.InfoLevel, zerolog.WarnLevel, zerolog.ErrorLevel,
		),
		validation.String(cfg.Directory, "directory").Required(true),
	)
}

type DatabaseConfig struct {
	fx.Out

	Directory string `env:"DIRECTORY" yaml:"directory" name:"db_directory"`
}

func (cfg DatabaseConfig) Validate() error {
	return validation.All(
		validation.String(cfg.Directory, "directory").Required(true),
	)
}

type HTTPConfig struct {
	fx.Out

	Port int `env:"PORT" yaml:"port" name:"http_port"`
}

func (cfg HTTPConfig) Validate() error {
	return validation.All(
		validation.Number(cfg.Port, "port").Required(true).With(isint.Port),
	)
}

func NewConfig(configPath string) (cfg Config, err error) {
	mode, ok := os.LookupEnv("APP_MODE")
	if !ok {
		mode = DebugMode
	} else if !ValidMode(mode) {
		return Config{}, fmt.Errorf(
			`invalid application mode "%s", expected one of "%s", "%s", "%s"`,
			mode, ReleaseMode, DebugMode, TestMode,
		)
	}

	SetMode(mode)

	if configPath != "" {
		var configs struct {
			Release Config `yaml:"release"`
			Debug   Config `yaml:"debug"`
			Test    Config `yaml:"test"`
		}

		err = cleanenv.ReadConfig(configPath, &configs)
		if err != nil {
			return Config{}, fmt.Errorf("could not read config: %w", err)
		}

		switch mode {
		case ReleaseMode:
			cfg = configs.Release
		case DebugMode:
			cfg = configs.Debug
		case TestMode:
			cfg = configs.Test
		}
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("could not read envvars: %w", err)
	}

	err = cfg.Validate()
	if err != nil {
		return Config{}, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}
