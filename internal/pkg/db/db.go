package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/ferdiebergado/lovemyride/internal/pkg/logger"
	"github.com/ferdiebergado/lovemyride/internal/pkg/parsers"
)

const (
	connMaxLifeTime    = 0
	maxIdleConnections = 50
	maxOpenConnections = 50
)

func Connect(ctx context.Context, getenv func(string) string) *sql.DB {
	db, err := sql.Open("pgx", getenv("DATABASE_URL"))

	if err != nil {
		logger.Fatal("Unable to connect to database", err)
	}

	db.SetConnMaxLifetime(time.Duration(parsers.ParseInt(getenv("DB_CONN_MAX_LIFETIME"), connMaxLifeTime)))
	db.SetMaxIdleConns(parsers.ParseInt(getenv("DB_MAX_IDLE_CONNECTIONS"), maxIdleConnections))
	db.SetMaxOpenConns(parsers.ParseInt(getenv("DB_MAX_OPEN_CONNECTIONS"), maxOpenConnections))

	return db
}
