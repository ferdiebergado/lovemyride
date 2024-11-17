package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
	"github.com/ferdiebergado/lovemyride/internal/pkg/logger"
)

type DeleteMode int

const (
	SoftDelete DeleteMode = iota
	HardDelete
)

var ErrRowClose = errors.New("failed to close the rows result set")
var ErrRowScan = errors.New("error occurred while scanning the row into the destination variables")
var ErrRowIteration = errors.New("error encountered during row iteration, possibly due to a database or connection issue")
var ErrModelNotFound = errors.New("model not found")

func Connect(ctx context.Context, dbConfig *config.DBOptions) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DB)

	db, err := sql.Open(dbConfig.Driver, dsn)

	if err != nil {
		logger.Fatal("open database", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		logger.Fatal("ping database", err)
	}

	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime))
	db.SetMaxIdleConns(dbConfig.MaxIdleConnections)
	db.SetMaxOpenConns(dbConfig.MaxOpenConnections)

	return db
}
