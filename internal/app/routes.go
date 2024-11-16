package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/go-express/middleware"
	"github.com/ferdiebergado/lovemyride/internal/app/spareparts"
	"github.com/ferdiebergado/lovemyride/internal/web/html"
)

// Setup router with middlewares
func SetupRouter(db *sql.DB) *router.Router {
	r := router.NewRouter()

	r.Use(middleware.RequestLogger)
	r.Use(middleware.PanicRecovery)

	AddRoutes(r, db)

	spareparts.AddRoutes(r, db)

	return r
}

func AddRoutes(router *router.Router, db *sql.DB) *router.Router {
	// Add routes here, see https://github.com/ferdiebergado/go-express for the documentation.
	router.Get("/home", func(w http.ResponseWriter, _ *http.Request) {
		html.Render(w, nil, "pages/home.html")
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		err := db.PingContext(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("Database connection failed: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("OK"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		html.Render(w, nil, "pages/404.html")
	})

	return router
}
