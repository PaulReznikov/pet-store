package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
  SERVER_PORT string
  DB_HOST string
  DB_PORT string
  DB_USER string
  DB_PASSWORD string
  DB_NAME string
  DB_SSLMODE string
}

func LoadConfig() (*Config, error) {
  if err := godotenv.Load(); err != nil {
    return nil, err
  }

  cfg := &Config{
    SERVER_PORT: os.Getenv("SERVER_PORT"),
    DB_HOST: os.Getenv("DB_HOST"),
    DB_PORT: os.Getenv("DB_PORT"),
    DB_USER: os.Getenv("DB_USER"),
    DB_PASSWORD: os.Getenv("DB_PASSWORD"),
    DB_NAME: os.Getenv("DB_NAME"),
    DB_SSLMODE: os.Getenv("DB_SSLMODE"),
  }
  if err := ValidateConfig(cfg); err != nil {
    return nil, err
  }
  return cfg, nil 
}

func ValidateConfig(cfg *Config) error {
  switch {
  default: return  nil
  case cfg.SERVER_PORT == "":
    return fmt.Errorf("SERVER_PORT is required")
  case cfg.DB_HOST == "":
    return fmt.Errorf("DB_HOST is required")
  case cfg.DB_PORT == "":
    return fmt.Errorf("DB_PORT is required")
  case cfg.DB_USER == "":
    return fmt.Errorf("DB_USER is required")
  case cfg.DB_PASSWORD == "":
    return fmt.Errorf("DB_PASSWORD is required")
  case cfg.DB_NAME == "":
    return fmt.Errorf("DB_NAME is required")
  case cfg.DB_SSLMODE == "":
    return fmt.Errorf("DB_SSLMODE is required")
  }
}