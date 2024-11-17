package spareparts

import (
	router "github.com/ferdiebergado/go-express"
)

func AddRoutes(router *router.Router, handler Handler) {
	router.Post("/spareparts", handler.CreateSparePart)
	router.Get("/spareparts/{id}", handler.GetSparePart)
	router.Get("/spareparts", handler.GetAllSpareParts)
	router.Patch("/spareparts/{id}", handler.UpdateSparePart)
	router.Delete("/spareparts/{id}", handler.DeleteSparePart)
}
