package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ngoldack/vercel-chi/api/handler"
)

func (app *App) NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	rh := handler.NewRootHandler(app.db)

	r.Get("/", rh.IndexHandler())

	return r
}
