package spareparts

import (
	"database/sql"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
)

func Mount(r *router.Router, db *sql.DB, config *config.Config) {
	repo := NewSparePartRepo(db)
	service := NewSparePartService(repo)
	handler := NewSparePartsHandler(service)

	AddRoutes(r, *handler, config)
}
