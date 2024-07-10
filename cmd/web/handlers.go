package main

import (
	"fmt"
	"html/template" // New import
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("in home", "method", "GET", "path", "/")

	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		logger.Info(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Info(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
