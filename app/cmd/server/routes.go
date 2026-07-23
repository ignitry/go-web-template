package main

import (
	"net/http"

	"starter/internal/page"
	"starter/internal/respond"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", health)
	mux.HandleFunc("GET /{$}", page.Home)
	mux.Handle("GET /assets/", page.Static())
	mux.HandleFunc("GET /404", respond.NotFound)
	mux.HandleFunc("/", notFoundRedirect)
	return mux
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok\\n"))
}

func notFoundRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/404", http.StatusFound)
}
