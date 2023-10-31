package api

import (
	"net/http"

	"github.com/ngoldack/vercel-chi/pkg/router"
)

// Handler is the entrypoint for the vercel serverless function
func Handler(w http.ResponseWriter, req *http.Request) {
	r := router.NewRouter(req.Context())
	r.ServeHTTP(w, req)
}
