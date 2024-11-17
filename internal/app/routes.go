package app

import (
	router "github.com/ferdiebergado/go-express"
)

func AddRoutes(router *router.Router, handler Handler) {
	router.Get("/home", handler.RenderHome)
	router.Get("/health", handler.CheckHealth)
	router.NotFound(handler.HandleNotFound)
}
