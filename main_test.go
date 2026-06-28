package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp map[string]string
	json.NewDecoder(w.Body).Decode(&resp)
	if resp["status"] != "ok" {
		t.Errorf("expected status ok, got %s", resp["status"])
	}
	if resp["service"] != "products-go" {
		t.Errorf("expected service products-go, got %s", resp["service"])
	}
}

func TestProductsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	productsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp []Product
	json.NewDecoder(w.Body).Decode(&resp)
	if len(resp) == 0 {
		t.Error("expected products, got empty list")
	}
	if resp[0].ID == 0 || resp[0].Name == "" {
		t.Error("product missing required fields")
	}
}

func TestProductsCount(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	productsHandler(w, req)

	var resp []Product
	json.NewDecoder(w.Body).Decode(&resp)
	if len(resp) != 5 {
		t.Errorf("expected 5 products, got %d", len(resp))
	}
}
