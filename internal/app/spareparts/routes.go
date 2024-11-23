package spareparts

import (
	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
)

func AddRoutes(router *router.Router, handler Handler, config *config.Config) {
	path := "/spareparts"
	pathID := path + "/" + config.App.IDPathValue
	resource := config.App.APIPrefix + path
	resourceID := resource + "/" + config.App.IDPathValue

	router.Post(resource, handler.CreateSparePart)
	router.Get(resourceID, handler.GetSparePart)
	router.Get(resource, handler.GetAllSpareParts)
	router.Patch(resourceID, handler.UpdateSparePart)
	router.Delete(resourceID, handler.DeleteSparePart)
	router.Get(path, handler.ListSpareParts)
	router.Get(path+"/create", handler.ShowCreateForm)
	router.Get(pathID, handler.ViewSparePart)
	router.Get(pathID+"/edit", handler.EditSparePart)
}
