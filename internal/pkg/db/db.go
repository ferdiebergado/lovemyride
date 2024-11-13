package db

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/lovemyride/internal/pkg/logger"
)

const (
	dbDriver           = "pgx"
	connMaxLifeTime    = 0
	maxIdleConnections = 50
	maxOpenConnections = 50
)

func Connect(_ context.Context, dsn string) *sql.DB {
	db, err := sql.Open(dbDriver, dsn)

	if err != nil {
		logger.Fatal("Unable to connect to database", err)
	}

	db.SetConnMaxLifetime(connMaxLifeTime)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)

	return db
}
