package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

func (ih *RootHandler) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "Hello, World!"})
	}
}
