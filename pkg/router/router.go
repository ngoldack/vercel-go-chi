package router

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ngoldack/vercel-chi/pkg/handler"
)

func NewRouter(ctx context.Context) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	ih := handler.NewIndexHandler()
	r.Get("/", ih.IndexHandle())

	return r
}
