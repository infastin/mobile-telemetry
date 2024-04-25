package app

import (
	"fmt"
	"os"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/fx"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	fx.Out

	Logger   LoggerConfig   `env-prefix:"LOGGER_" yaml:"logger"`
	Database DatabaseConfig `env-prefix:"DATABASE_" yaml:"database"`
	HTTP     HTTPConfig     `env-prefix:"HTTP_" yaml:"http"`
}

func (cfg Config) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Logger),
		validation.Field(&cfg.Database),
		validation.Field(&cfg.HTTP),
	)
}

type LoggerConfig struct {
	fx.Out

	Level zapcore.Level `env:"LEVEL" yaml:"level" name:"logger_level"`
}

func (cfg LoggerConfig) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Level,
			validation.In(zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel),
		),
	)
}

type DatabaseConfig struct {
	fx.Out

	Directory string `env:"DIRECTORY" yaml:"directory" name:"db_directory"`
}

func (cfg DatabaseConfig) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Directory, validation.Required),
	)
}

type HTTPConfig struct {
	fx.Out

	Port int `env:"PORT" yaml:"port" name:"http_port"`
}

func (cfg HTTPConfig) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Port,
			validation.Required,
			validation.Min(0).Exclusive(),
			validation.Max(65536).Exclusive(),
		),
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

	validation.ErrorTag = "yaml"
	defer func() {
		validation.ErrorTag = "json"
	}()

	err = cfg.Validate()
	if err != nil {
		return Config{}, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}
