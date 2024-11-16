package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
	"github.com/ferdiebergado/lovemyride/internal/pkg/logger"
)

var ErrRowClose = errors.New("failed to close the rows result set")
var ErrRowScan = errors.New("error occurred while scanning the row into the destination variables")
var ErrRowIteration = errors.New("error encountered during row iteration, possibly due to a database or connection issue")
var ErrModelNotFound = errors.New("model not found")

func Connect(ctx context.Context, config *config.DBOptions) *sql.DB {
	db, err := sql.Open(config.Driver, config.DSN)

	if err != nil {
		logger.Fatal("open database", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		logger.Fatal("ping database", err)
	}

	db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime))
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetMaxOpenConns(config.MaxOpenConnections)

	return db
}
