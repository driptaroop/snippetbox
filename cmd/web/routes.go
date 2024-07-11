package main

import "net/http"

func (app *Application) routes(staticDir string) *http.ServeMux {
	fileServer := http.FileServer(http.Dir(staticDir))
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	return mux
}
