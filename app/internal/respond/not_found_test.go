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
	req.Header.Set("Accept", "text/html; q=0.5, application/json; q=0.9")
	res := httptest.NewRecorder()

	NotFound(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusNotFound)
	}
	if got := res.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("content type = %q", got)
	}
	if got := res.Header().Get("Vary"); got != "Accept" {
		t.Fatalf("vary = %q", got)
	}
	if got := res.Body.String(); got != "{\"error\":\"not found\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

func TestNotFoundAcceptQuality(t *testing.T) {
	for _, tc := range []struct {
		accept      string
		contentType string
	}{
		{"text/html, application/json;q=0.9", "text/html; charset=utf-8"},
		{"application/json;q=0, */*;q=1", "text/html; charset=utf-8"},
		{"text/html;q=0, application/*;q=0.5", "application/json; charset=utf-8"},
		{"*/*;q=0", ""},
	} {
		t.Run(tc.accept, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/missing", nil)
			req.Header.Set("Accept", tc.accept)
			res := httptest.NewRecorder()
			NotFound(res, req)

			if res.Code != http.StatusNotFound {
				t.Fatalf("status = %d, want %d", res.Code, http.StatusNotFound)
			}
			if got := res.Header().Get("Content-Type"); got != tc.contentType {
				t.Fatalf("content type = %q, want %q", got, tc.contentType)
			}
		})
	}
}
