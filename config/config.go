package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App      `yaml:"app"`
		Server   `yaml:"server"`
		Log      `yaml:"logger"`
		Customer `yaml:"customer"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	Server struct {
		Port     string `env-required:"true" yaml:"port" env:"TCP_PORT"`
		Protocol string `env-required:"true" yaml:"protocol" env:"PROTOCOL"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Customer struct {
		PrivateKeyPath string `env-required:"true" yaml:"privateKeyPath"    env:"PRIVATE_KEY_PATH"`
		User           string `env-required:"true" yaml:"user" env:"USER_SSH"`
		Host           string `env-required:"true" yaml:"host" env:"HOST"`
		SSHPort        string `env-required:"true" yaml:"ssh_port" env:"SSH_PORT"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
