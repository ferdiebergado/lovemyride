package app

import (
	"database/sql"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/go-express/middleware"
	"github.com/ferdiebergado/lovemyride/internal/app/spareparts"
)

// Setup router with middlewares
func SetupRouter(db *sql.DB) *router.Router {
	r := router.NewRouter()

	r.Use(middleware.RequestLogger)
	r.Use(middleware.PanicRecovery)

	handler := NewAppHandler(db)

	AddRoutes(r, *handler)

	spareparts.Mount(r, db)

	return r
}
