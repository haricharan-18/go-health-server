package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_Status200(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Code != 200 {
		t.Errorf("got %d, want 200", w.Code)
	}
}

func TestHealthHandler_ContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("wrong content-type: %s", w.Header().Get("Content-Type"))
	}
}

func TestHealthHandler_BodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Body.Len() == 0 {
		t.Error("body should not be empty")
	}
}

func TestHealthHandler_404OnUnknownPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	mux.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("got %d, want 404", w.Code)
	}
}

func TestHealthHandler_405OnPost(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Code != 405 {
		t.Errorf("got %d, want 405", w.Code)
	}
}