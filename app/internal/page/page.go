package page

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates/*.html views/*.html static/*
var files embed.FS

var (
	templates = template.Must(template.ParseFS(files, "templates/application.html", "views/home.html"))
	staticFS  = mustSub(files, "static")
)

func Home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = templates.ExecuteTemplate(w, "application.html", struct{ Title string }{Title: "Starter"})
}

func Static() http.Handler {
	return http.StripPrefix("/assets/", http.FileServer(http.FS(staticFS)))
}

func mustSub(files fs.FS, directory string) fs.FS {
	sub, err := fs.Sub(files, directory)
	if err != nil {
		panic(err)
	}
	return sub
}
