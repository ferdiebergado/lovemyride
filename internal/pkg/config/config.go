package config

import (
	"os"

	"github.com/ferdiebergado/lovemyride/internal/pkg/env"
)

type Config struct {
	App    *AppOptions
	Server *ServerOptions
	DB     *DBOptions
}

type AppOptions struct {
	APIPrefix   string
	IDPathValue string
}

type ServerOptions struct {
	Host            string
	Port            string
	ReadTimeout     int
	WriteTimeout    int
	IdleTimeout     int
	ShutdownTimeout int
}

type DBOptions struct {
	Driver             string
	Host               string
	Port               string
	User               string
	Password           string
	DB                 string
	ConnMaxLifetime    int
	MaxIdleConnections int
	MaxOpenConnections int
}

const (
	serverShutdownTimeout = 10
	serverReadTimeout     = 10
	serverWriteTimeout    = 10
	serverIdleTimeout     = 60

	connMaxLifeTime    = 0
	maxIdleConnections = 50
	maxOpenConnections = 50
)

func NewAppConfig() *Config {
	return &Config{
		App: &AppOptions{
			APIPrefix:   "/api",
			IDPathValue: "{id}",
		},
		Server: &ServerOptions{
			Host:            os.Getenv("HOST"),
			Port:            env.GetEnv("PORT", "8000"),
			ReadTimeout:     serverReadTimeout,
			WriteTimeout:    serverWriteTimeout,
			IdleTimeout:     serverIdleTimeout,
			ShutdownTimeout: serverShutdownTimeout,
		},
		DB: &DBOptions{
			Driver:             "pgx",
			Host:               env.Must("DB_HOST"),
			Port:               env.Must("DB_PORT"),
			User:               env.Must("DB_USER"),
			Password:           env.Must("DB_PASS"),
			DB:                 env.Must("DB_NAME"),
			ConnMaxLifetime:    connMaxLifeTime,
			MaxIdleConnections: maxIdleConnections,
			MaxOpenConnections: maxOpenConnections,
		},
	}
}
