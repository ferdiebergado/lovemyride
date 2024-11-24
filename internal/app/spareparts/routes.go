package spareparts

import (
	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
)

func AddRoutes(router *router.Router, handler Handler, config *config.Config) {
	const path = "/spareparts"
	idPath := "/{" + config.App.IDPathValue + "}"
	pathWithID := path + idPath
	resource := config.App.APIPrefix + path
	resourceWithID := resource + idPath

	// API routes
	router.Post(resource, handler.CreateSparePart)
	router.Get(resourceWithID, handler.GetSparePart)
	router.Get(resource, handler.GetAllSpareParts)
	router.Patch(resourceWithID, handler.UpdateSparePart)
	router.Delete(resourceWithID, handler.DeleteSparePart)

	// HTML routes
	router.Get(path, handler.ListSpareParts)
	router.Get(path+"/create", handler.ShowCreateForm)
	router.Get(pathWithID, handler.ViewSparePart)
	router.Get(pathWithID+"/edit", handler.EditSparePart)
}
