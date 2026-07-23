package respond

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNotFoundHTML(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	res := httptest.NewRecorder()

	NotFound(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusNotFound)
	}
	if got := res.Header().Get("Content-Type"); got != "text/html; charset=utf-8" {
		t.Fatalf("content type = %q", got)
	}
	if !strings.Contains(res.Body.String(), "<h1>404</h1>") {
		t.Fatalf("body = %q", res.Body.String())
	}
}

func TestNotFoundJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	req.Header.Set("Accept", "text/html, application/json; q=0.9")
	res := httptest.NewRecorder()

	NotFound(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusNotFound)
	}
	if got := res.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("content type = %q", got)
	}
	if got := res.Body.String(); got != "{\"error\":\"not found\"}\n" {
		t.Fatalf("body = %q", got)
	}
}
