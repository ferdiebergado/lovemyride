package spareparts

import (
	"database/sql"

	router "github.com/ferdiebergado/go-express"
)

func Mount(r *router.Router, db *sql.DB) {
	repo := NewSparePartRepo(db)
	service := NewSparePartService(repo)
	handler := NewSparePartsHandler(service)

	AddRoutes(r, *handler)
}
