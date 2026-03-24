package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SERVER_HOST string `required:"true" envconfig:"SERVER_HOST"`
	SERVER_PORT string `required:"true" envconfig:"SERVER_PORT"`
	DB_HOST     string `required:"true" envconfig:"DB_HOST"`
	DB_PORT     string `required:"true" envconfig:"DB_PORT"`
	DB_USER     string `required:"true" envconfig:"DB_USER"`
	DB_PASSWORD string `required:"true" envconfig:"DB_PASSWORD"`
	DB_NAME     string `required:"true" envconfig:"DB_NAME"`
	DB_SSLMODE  string `required:"true" envconfig:"DB_SSLMODE"`
}

func LoadConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) CreateDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.DB_USER, c.DB_PASSWORD, c.DB_HOST, c.DB_PORT, c.DB_NAME, c.DB_SSLMODE)
}
