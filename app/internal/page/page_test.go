package page

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeRendersApplicationAndHomeTemplates(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	Home(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	for _, text := range []string{"<title>Starter</title>", "Replace this page.", "/assets/app.css"} {
		if !strings.Contains(response.Body.String(), text) {
			t.Fatalf("body does not contain %q", text)
		}
	}
}

func TestStaticServesCSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/assets/app.css", nil)
	response := httptest.NewRecorder()

	Static().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if !strings.Contains(response.Header().Get("Content-Type"), "text/css") {
		t.Fatalf("content type = %q", response.Header().Get("Content-Type"))
	}
}
