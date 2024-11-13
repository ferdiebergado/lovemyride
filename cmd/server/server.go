package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/lovemyride/internal/pkg/parsers"
)

const (
	serverShutdownTimeout = 10
	serverReadTimeout     = 10
	serverWriteTimeout    = 10
	serverIdleTimeout     = 60
)

func createServer(getenv func(string) string, router *router.Router) *http.Server {
	// Parse server timeouts
	readTimeout := parsers.ParseTimeout(getenv("SERVER_READ_TIMEOUT"), serverReadTimeout)
	writeTimeout := parsers.ParseTimeout(getenv("SERVER_WRITE_TIMEOUT"), serverWriteTimeout)
	idleTimeout := parsers.ParseTimeout(getenv("SERVER_IDLE_TIMEOUT"), serverIdleTimeout)

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
