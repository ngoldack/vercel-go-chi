package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ngoldack/vercel-chi/pkg/router"
)

// This file is only used for local development
func main() {
	ctx := context.Background()
	slog.WarnContext(ctx, "DO NOT USE THIS SERVER IN PRODUCTION!")

	r := router.NewRouter(ctx)
	slog.InfoContext(ctx, "Local dev server started!", "addr", "http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}
