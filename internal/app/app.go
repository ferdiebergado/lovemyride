package app

import (
	"database/sql"
	"log/slog"

	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/go-express/middleware"
	"github.com/ferdiebergado/lovemyride/internal/app/spareparts"
	"github.com/ferdiebergado/lovemyride/internal/pkg/config"
)

type App struct {
	DB     *sql.DB
	Config *config.Config
	Logger *slog.Logger
	Router *router.Router
}

func NewApp(db *sql.DB, config *config.Config, logger *slog.Logger) *App {
	app := &App{
		DB:     db,
		Config: config,
		Logger: logger,
		Router: router.NewRouter(),
	}

	app.SetupRoutes()

	return app
}

// Setup router with middlewares
func (a *App) SetupRoutes() {
	// Global middlewares
	a.Router.Use(middleware.RequestLogger)
	a.Router.Use(middleware.PanicRecovery)

	// Base routes
	handler := NewAppHandler(a.DB)
	a.Router.Get("/home", handler.RenderHome)
	a.Router.Get("/health", handler.CheckHealth)
	a.Router.NotFound(handler.HandleNotFound)

	// Spareparts routes
	spareparts.Mount(a.Router, a.DB, a.Config)
}
