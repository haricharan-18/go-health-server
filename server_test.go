package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_Returns200(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestHealthHandler_ContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	ct := w.Header().Get("Content-Type")

	if ct != "application/json" {
		t.Errorf("expected application/json, got %s", ct)
	}
}

func TestHealthHandler_BodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Body.Len() == 0 {
		t.Error("expected non-empty body")
	}
}

func TestHealthHandler_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", w.Code)
	}
}
