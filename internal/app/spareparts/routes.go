package spareparts

import (
	"database/sql"

	router "github.com/ferdiebergado/go-express"
)

func AddRoutes(router *router.Router, db *sql.DB) *router.Router {
	repo := NewSparePartRepo(db)
	service := NewSparePartService(repo)
	handler := NewSparePartsHandler(service)

	router.Post("/spareparts", handler.CreateSparePart)
	router.Get("/spareparts/{id}", handler.GetSpareParts)
	router.Get("/spareparts", handler.GetAllSpareParts)
	router.Put("/spareparts/{id}", handler.UpdateSparePart)
	router.Delete("/spareparts/{id}", handler.DeleteSparePart)

	return router
}
