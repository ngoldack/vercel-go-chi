package main

import (
	"log"
	"net/http"

	"github.com/ngoldack/vercel-chi/internal/app"
)

// Vercel is the entrypoint for the vercel serverless function
// It is defined in api/vercel.go
// This file is only used for local development
func main() {
	app := app.NewApp()
	r := app.NewRouter()

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", r)
}
