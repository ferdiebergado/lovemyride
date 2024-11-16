package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/lovemyride/internal/app"
	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
	"github.com/ferdiebergado/lovemyride/internal/pkg/parsers"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	serverShutdownTimeout = 10
	serverReadTimeout     = 10
	serverWriteTimeout    = 10
	serverIdleTimeout     = 60
)

func setupLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func createServer(getenv func(string) string, router *router.Router) *http.Server {
	// Parse server timeouts
	readTimeout := parsers.ParseInt(getenv("SERVER_READ_TIMEOUT"), serverReadTimeout)
	writeTimeout := parsers.ParseInt(getenv("SERVER_WRITE_TIMEOUT"), serverWriteTimeout)
	idleTimeout := parsers.ParseInt(getenv("SERVER_IDLE_TIMEOUT"), serverIdleTimeout)

	// Configure HTTP server
	return &http.Server{
		Addr:         ":" + getenv("PORT"),
		Handler:      router,
		ReadTimeout:  time.Duration(readTimeout * int(time.Second)),
		WriteTimeout: time.Duration(writeTimeout * int(time.Second)),
		IdleTimeout:  time.Duration(idleTimeout * int(time.Second)),
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
func waitForShutdown(ctx context.Context, httpServer *http.Server, stderr io.Writer) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
}

func run(ctx context.Context, _ []string, getenv func(string) string, _ io.Reader, _, stderr io.Writer) error {
	// Handle OS interrupt signals
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// Initialize logger
	setupLogger()

	// Connect to the database
	conn := db.Connect(ctx, getenv)
	defer conn.Close()

	// Initialize router and add middlewares
	router := app.SetupRouter(conn)

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
