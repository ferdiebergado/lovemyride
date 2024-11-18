package spareparts

import (
	router "github.com/ferdiebergado/go-express"
)

func AddRoutes(router *router.Router, handler Handler) {
	router.Post("/api/spareparts", handler.CreateSparePart)
	router.Get("/api/spareparts/{id}", handler.GetSparePart)
	router.Get("/api/spareparts", handler.GetAllSpareParts)
	router.Patch("/api/spareparts/{id}", handler.UpdateSparePart)
	router.Delete("/api/spareparts/{id}", handler.DeleteSparePart)
	router.Get("/spareparts", handler.ListSpareParts)
	router.Get("/spareparts/create", handler.ShowCreateForm)
	router.Get("/spareparts/{id}", handler.ViewSparePart)
	router.Get("/spareparts/{id}/edit", handler.EditSparePart)
}
