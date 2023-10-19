package api

import (
	"net/http"

	"github.com/ngoldack/vercel-chi/internal/app"
)

// Handler is the entrypoint for the vercel serverless function
func Handler(w http.ResponseWriter, req *http.Request) {
	app := app.NewApp()
	r := app.NewRouter()

	r.ServeHTTP(w, req)
}
