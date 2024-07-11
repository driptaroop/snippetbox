package main

import "log/slog"

type Application struct {
	logger *slog.Logger
	config *Config
}

type Config struct {
	addr      string
	staticDir string
	dsn       string
}
