package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
	"github.com/ferdiebergado/lovemyride/internal/pkg/env"
	"github.com/ferdiebergado/lovemyride/internal/pkg/logging"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestApp(t *testing.T) {
	err := env.LoadEnv("../../.env.testing")

	if err != nil {
		t.Logf("load env: %v", err)
		t.FailNow()
	}

	appConfig := config.NewAppConfig()

	conn := db.Connect(context.Background(), appConfig.DB)
	defer conn.Close()

	logger := logging.CreateLogger()

	app := NewApp(conn, appConfig, logger)

	r := app.Router

	t.Run("GET / should return status 200 and render home.html", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/home", nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected %d but got %d", http.StatusOK, rec.Code)
		}

		// Check if the response contains content from home.html
		expected := "Welcome!"
		actual := rec.Body.String()

		if !strings.Contains(actual, expected) {
			t.Errorf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("GET /health should return status 200 and return OK", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected %d but got %d", http.StatusOK, rec.Code)
		}

		// Check if the response contains content from home.html
		expected := "OK"
		actual := rec.Body.String()

		if actual != expected {
			t.Errorf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("GET /nonexistent should return status 404 and render 404.html", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("Expected %d but got %d", http.StatusNotFound, rec.Code)
		}

		expected := "The page you are looking for does not exist."
		actual := rec.Body.String()

		// Check if the response contains content from 404.html
		if !strings.Contains(actual, expected) {
			t.Errorf("Expected %s but got %s", expected, actual)
		}
	})
}
