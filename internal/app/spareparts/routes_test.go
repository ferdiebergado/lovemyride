package spareparts

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var appConfig = config.NewAppConfig()

var conn = db.Connect(context.Background(), appConfig.DB)

func TestAddRoutes(t *testing.T) {
	r := router.NewRouter() // Create a new instance of your custom Router
	AddRoutes(r, conn)      // Add your routes to the Router

	t.Run("POST /spareparts should return status 201 and return json data of spareparts", func(t *testing.T) {
		payload := CreateParams{
			Description:         "Spark Plug",
			MaintenanceInterval: 8000,
		}

		// Convert the struct to JSON
		jsonBody, err := json.Marshal(payload)

		if err != nil {
			t.Fatalf("Failed to marshal JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/spareparts", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Errorf("Expected %d but got %d", http.StatusCreated, rec.Code)
		}
	})
}
