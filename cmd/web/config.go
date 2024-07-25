package main

import (
	"github.com/alexedwards/scs/v2" // New import
	"github.com/go-playground/form/v4"
	"html/template"
	"log/slog"
	"snippetbox.dripto.org/internal/models"
)

type Application struct {
	logger         *slog.Logger
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}
