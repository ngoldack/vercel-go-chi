package handler_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/ngoldack/vercel-chi/pkg/handler"
	"github.com/stretchr/testify/require"
)

func TestIndexHandle(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	rootHandler := handler.NewIndexHandler()
	rootHandler.IndexHandle().ServeHTTP(w, req)

	result := w.Result()
	require.Equal(t, 200, result.StatusCode, "should return 200 OK")
	require.Equal(t, "application/json", result.Header.Get("Content-Type"), "should return JSON")

	m := make(map[string]string)
	err := json.NewDecoder(w.Result().Body).Decode(&m)
	require.NoError(t, err, "should decode JSON")

	require.True(t, len(m["message"]) > 0, "should return a message")
}
