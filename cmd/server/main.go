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

	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
