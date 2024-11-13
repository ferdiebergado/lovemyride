package app

import (
	"net/http"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/go-express/middleware"
	"github.com/ferdiebergado/lovemyride/internal/web/html"
)

// Setup router with middlewares
func SetupRouter() *router.Router {
	r := router.NewRouter()
	r.Use(middleware.RequestLogger)
	r.Use(middleware.PanicRecovery)
	AddRoutes(r)
	return r
}

func AddRoutes(router *router.Router) *router.Router {
	// Add routes here, see https://github.com/ferdiebergado/go-express for the documentation.
	router.Get("/home", func(w http.ResponseWriter, _ *http.Request) {
		html.Render(w, nil, "pages/home.html")
	})

	router.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		html.Render(w, nil, "pages/404.html")
	})

	return router
}
