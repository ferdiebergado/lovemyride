package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
	"github.com/ferdiebergado/lovemyride/internal/pkg/logger"
)

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
