package respond

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
)

//go:embed templates/404.html
var notFoundHTML string

var notFoundTemplate = template.Must(template.New("404.html").Parse(notFoundHTML))

func NotFound(w http.ResponseWriter, r *http.Request) {
	if acceptsJSON(r.Header.Get("Accept")) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	_ = notFoundTemplate.Execute(w, nil)
}

func acceptsJSON(accept string) bool {
	for _, value := range strings.Split(accept, ",") {
		mediaType := strings.TrimSpace(strings.SplitN(value, ";", 2)[0])
		if mediaType == "application/json" {
			return true
		}
	}
	return false
}
