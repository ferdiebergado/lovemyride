package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"

	"github.com/ferdiebergado/lovemyride/internal/app"
	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func setupLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func run(ctx context.Context, _ []string, getenv func(string) string, _ io.Reader, _, stderr io.Writer) error {
	// Handle OS interrupt signals
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// Initialize logger
	setupLogger()

	// Connect to the database
	conn := db.Connect(ctx, getenv("DATABASE_URL"))
	defer conn.Close()

	// Initialize router and add middlewares
	router := app.SetupRouter()

	// Create the server
	httpServer := createServer(getenv, router)

	// Start the server in a goroutine
	go startServer(httpServer, stderr)

	// Wait for shutdown
	waitForShutdown(ctx, httpServer, stderr)

	return nil
}

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
