package config

import "errors"

// ServerConfig contains configuration for server to run.
type ServerConfig struct {
	APIPort    int    `env:"API_PORT" envDefault:"9000"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
}

var (
	// ErrDbUsernameEmpty is returned if DB username is not provided.
	ErrDbUsernameEmpty = errors.New("DB_USERNAME is empty")
	// ErrDbPasswordEmpty is returned if DB password is not provided.
	ErrDbPasswordEmpty = errors.New("DB_PASSWORD is empty")
	// ErrDbNameEmpty is returned if DB name is not provided.
	ErrDbNameEmpty = errors.New("DB_NAME is empty")
)
