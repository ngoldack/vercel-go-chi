package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (ih *IndexHandler) IndexHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]any{"message": "Hello, World!"})
	}
}
