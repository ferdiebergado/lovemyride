package spareparts

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
	"github.com/ferdiebergado/lovemyride/internal/pkg/env"
	"github.com/ferdiebergado/lovemyride/internal/pkg/http/response"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var conn *sql.DB
var r *router.Router
var repo Repo

func setup() {
	err := env.LoadEnv("../../../.env.testing")

	if err != nil {
		log.Fatal("Can't load .env file")
	}

	appConfig := config.NewAppConfig()
	conn = db.Connect(context.Background(), appConfig.DB)
	repo = NewSparePartRepo(conn)
	r = router.NewRouter()

	Mount(r, conn)
}

func createSparePart(ctx context.Context, params CreateParams) (SparePart, error) {
	return repo.Create(ctx, params)
}

func getSparePart(ctx context.Context, id string) (SparePart, error) {
	return repo.Find(ctx, id)
}

func tearDown() {
	_, err := conn.ExecContext(context.Background(), `DELETE FROM spareparts`)

	if err != nil {
		log.Fatalf("delete spareparts: %+v", err)
	}
}

func TestSpareParts(t *testing.T) {
	setup()

	t.Run("POST /spareparts should return status 201 and return json data of the created sparepart", func(t *testing.T) {
		payload := CreateParams{
			Description:         "Spark Plug",
			MaintenanceInterval: 8000,
		}

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

		var apiResponse response.APIResponse[SparePart]

		err = json.NewDecoder(rec.Body).Decode(&apiResponse)

		if err != nil {
			t.Fatalf("Failed to decode JSON: %v\n", err)
		}

		sparePart := apiResponse.Data

		if sparePart.ID == "" || sparePart.CreatedAt.IsZero() || sparePart.UpdatedAt.IsZero() {
			t.Errorf("Unexpected values: %+v", sparePart)
		}

		if sparePart.Description != payload.Description || sparePart.MaintenanceInterval != payload.MaintenanceInterval {
			t.Errorf("Unexpected values: %+v", sparePart)
		}

		tearDown()
	})

	t.Run("GET /spareparts/{id} should return status 200 and return json data of the specified sparepart", func(t *testing.T) {
		payload := CreateParams{
			Description:         "ECU",
			MaintenanceInterval: 100000,
		}

		ctx := context.Background()

		newSparePart, err := createSparePart(ctx, payload)

		if err != nil {
			t.Error("Unable to create spare part")
		}

		ID := newSparePart.ID

		req := httptest.NewRequest(http.MethodGet, "/spareparts/"+ID, nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected %d but got %d", http.StatusOK, rec.Code)
		}

		var apiResponse response.APIResponse[SparePart]

		err = json.NewDecoder(rec.Body).Decode(&apiResponse)

		if err != nil {
			t.Fatalf("Failed to decode JSON: %v\n", err)
		}

		sparePart := apiResponse.Data

		if !reflect.DeepEqual(newSparePart, sparePart) {
			t.Errorf("Unexpected values: %+v", sparePart)
		}

		tearDown()
	})

	t.Run("GET /spareparts should return status 200 and return json data with slice of spareparts", func(t *testing.T) {
		items := []CreateParams{
			{
				Description:         "ECU",
				MaintenanceInterval: 100000,
			},
			{
				Description:         "Engine Oil",
				MaintenanceInterval: 2000,
			},
			{
				Description:         "Drive Belt",
				MaintenanceInterval: 24000,
			},
		}

		var spareParts []SparePart

		ctx := context.Background()

		for _, item := range items {
			sparePart, err := createSparePart(ctx, item)
			if err != nil {
				t.Error("Unable to create sparepart")
			}
			spareParts = append(spareParts, sparePart)
		}

		req := httptest.NewRequest(http.MethodGet, "/spareparts", nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		var apiResponse response.APIResponse[[]SparePart]

		err := json.NewDecoder(rec.Body).Decode(&apiResponse)

		if err != nil {
			t.Fatalf("Failed to decode JSON: %v\n", err)
		}

		data := apiResponse.Data

		if len(data) < len(spareParts) {
			t.Errorf("Expected >= %d but got %d", len(spareParts), len(data))
		}

		tearDown()
	})

	t.Run("PATCH /spareparts/{id} should return status 200 and return json data with success", func(t *testing.T) {
		payload := CreateParams{
			Description:         "ECU",
			MaintenanceInterval: 100000,
		}

		ctx := context.Background()

		newSparePart, err := createSparePart(ctx, payload)

		if err != nil {
			t.Error("Unable to create spare part")
		}

		ID := newSparePart.ID

		updateParams := UpdateParams{
			MaintenanceInterval: 200000,
		}

		jsonBody, err := json.Marshal(updateParams)

		if err != nil {
			t.Fatalf("Failed to marshal JSON: %v", err)
		}

		req := httptest.NewRequest(http.MethodPatch, "/spareparts/"+ID, bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected %d but got %d", http.StatusOK, rec.Code)
		}

		var apiResponse response.APIResponse[SparePart]

		err = json.NewDecoder(rec.Body).Decode(&apiResponse)

		if err != nil {
			t.Fatalf("Failed to decode JSON: %v\n", err)
		}

		if !apiResponse.Success {
			t.Errorf("Expected %t but got %t", true, apiResponse.Success)
		}

		sparePart, err := getSparePart(ctx, ID)

		if err != nil {
			t.Fatalf("get spare part: %+v", err)
		}

		if sparePart.MaintenanceInterval != updateParams.MaintenanceInterval {
			t.Errorf("Expected %d but got %d", updateParams.MaintenanceInterval, sparePart.MaintenanceInterval)
		}

		if newSparePart.Description != sparePart.Description {
			t.Errorf("Expected %s but got %s", newSparePart.Description, sparePart.Description)
		}

		tearDown()
	})

	t.Run("DELETE /spareparts/{id} should return status 200 and return json data with success", func(t *testing.T) {
		payload := CreateParams{
			Description:         "ECU",
			MaintenanceInterval: 100000,
		}

		ctx := context.Background()

		newSparePart, err := createSparePart(ctx, payload)

		if err != nil {
			t.Error("Unable to create spare part")
		}

		t.Logf("newsparepart: %+v", newSparePart)

		ID := newSparePart.ID

		req := httptest.NewRequest(http.MethodDelete, "/spareparts/"+ID, nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected %d but got %d", http.StatusOK, rec.Code)
		}

		var apiResponse response.APIResponse[SparePart]

		err = json.NewDecoder(rec.Body).Decode(&apiResponse)

		if err != nil {
			t.Fatalf("Failed to decode JSON: %v\n", err)
		}

		if !apiResponse.Success {
			t.Errorf("Expected %t but got %t", true, apiResponse.Success)
		}

		sparePart, err := getSparePart(ctx, ID)

		if err != nil {
			t.Fatalf("get spare part: %+v", err)
		}

		if !sparePart.DeletedAt.Valid {
			t.Errorf("Expected %t but got %t", true, sparePart.DeletedAt.Valid)
		}

		tearDown()
	})
}
