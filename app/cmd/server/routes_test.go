package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutesNotFound(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/missing", nil)
	response := httptest.NewRecorder()

	routes().ServeHTTP(response, request)

	if response.Code != http.StatusFound {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusFound)
	}
	if location := response.Header().Get("Location"); location != "/404" {
		t.Fatalf("location = %q, want %q", location, "/404")
	}
}
