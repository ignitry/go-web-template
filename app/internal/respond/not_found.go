package respond

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

//go:embed templates/404.html
var notFoundHTML string

var notFoundTemplate = template.Must(template.New("404.html").Parse(notFoundHTML))

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Vary", "Accept")
	accept := strings.Join(r.Header.Values("Accept"), ",")
	htmlQuality, htmlSpecificity := mediaQuality(accept, "text/html", "text/*")
	jsonQuality, jsonSpecificity := mediaQuality(accept, "application/json", "application/*")
	if jsonQuality > htmlQuality || (jsonQuality > 0 && jsonQuality == htmlQuality && jsonSpecificity > htmlSpecificity) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
		return
	}

	if htmlQuality == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	_ = notFoundTemplate.Execute(w, nil)
}

func mediaQuality(accept, exact, wildcard string) (quality float64, specificity int) {
	if strings.TrimSpace(accept) == "" {
		return 1, 0
	}

	specificity = -1
	for _, value := range strings.Split(accept, ",") {
		mediaType, params, err := mime.ParseMediaType(value)
		if err != nil {
			continue
		}

		var match int
		switch mediaType {
		case exact:
			match = 2
		case wildcard:
			match = 1
		case "*/*":
			match = 0
		default:
			continue
		}

		q := 1.0
		if raw, ok := params["q"]; ok {
			q, err = strconv.ParseFloat(raw, 64)
			if err != nil || !(q >= 0 && q <= 1) {
				continue
			}
		}

		if match > specificity || (match == specificity && q > quality) {
			quality, specificity = q, match
		}
	}
	return quality, specificity
}
