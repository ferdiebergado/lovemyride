package app

import (
	"database/sql"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/go-express/middleware"
	"github.com/ferdiebergado/lovemyride/internal/app/spareparts"
	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
)

// Setup router with middlewares
func SetupRouter(db *sql.DB, config *config.Config) *router.Router {
	r := router.NewRouter()

	r.Use(middleware.RequestLogger)
	r.Use(middleware.PanicRecovery)

	handler := NewAppHandler(db)

	AddRoutes(r, *handler)

	spareparts.Mount(r, db, config)

	return r
}
