package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/lovemyride/internal/app"
	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
	"github.com/ferdiebergado/lovemyride/internal/pkg/env"
	"github.com/ferdiebergado/lovemyride/internal/pkg/logging"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func createServer(config *config.ServerOptions, router *router.Router) *http.Server {
	// Configure HTTP server
	return &http.Server{
		Addr:         ":" + config.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.IdleTimeout) * time.Second,
	}
}

// Start the HTTP server and handle errors
func startServer(httpServer *http.Server, stderr io.Writer) {
	fmt.Printf("HTTP Server listening on %s...\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintf(stderr, "error listening and serving: %s\n", err)
	}
}

// Wait for context cancellation and gracefully shut down the server
func waitForShutdown(ctx context.Context, httpServer *http.Server, config *config.ServerOptions, stderr io.Writer) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ShutdownTimeout)*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
}

func run(ctx context.Context, _ []string, _ io.Reader, _, stderr io.Writer) error {
	// Handle OS interrupt signals
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// Initialize app config
	config := config.NewAppConfig()

	// Connect to the database
	conn := db.Connect(ctx, config.DB)
	defer conn.Close()

	// Create a logger
	logger := logging.CreateLogger()

	// Setup a new app
	app := app.NewApp(conn, config, logger)

	// Create the server
	httpServer := createServer(config.Server, app.Router)

	// Start the server in a goroutine
	go startServer(httpServer, stderr)

	// Wait for shutdown
	waitForShutdown(ctx, httpServer, config.Server, stderr)

	return nil
}

func main() {
	appEnv := env.GetEnv("GO_ENV", "development")

	if appEnv != "production" {
		err := env.LoadEnv(".env." + appEnv)

		if err != nil {
			logging.Fatal("Unable to load .env file", err)
		}
	}

	if err := run(context.Background(), os.Args, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
