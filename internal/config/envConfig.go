package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"os"
)

// FromEnvVariables function populates ServerConfig from environment variables.
func FromEnvVariables() (*ServerConfig, error) {
	environment := "development"

	if val, exists := os.LookupEnv("APP_ENV"); exists {
		environment = val
	}

	err := godotenv.Overload(".env." + environment)
	if err != nil {
		return nil, err
	}

	var cfg ServerConfig
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	if cfg.DBUsername == "" {
		return nil, ErrDbUsernameEmpty
	}

	if cfg.DBPassword == "" {
		return nil, ErrDbPasswordEmpty
	}

	if cfg.DBName == "" {
		return nil, ErrDbNameEmpty
	}

	return &cfg, nil
}
