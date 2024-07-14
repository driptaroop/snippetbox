package main

import (
	"html/template"
	"log/slog"
	"snippetbox.dripto.org/internal/models"
)

type Application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}
