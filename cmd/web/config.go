package main

import (
	"log/slog"
	"snippetbox.dripto.org/internal/models"
)

type Application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}
