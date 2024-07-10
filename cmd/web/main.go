package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

func main() {
	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "staticDir", "./ui/static/", "static directory")
	flag.Parse()

	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	logger.Info(fmt.Sprintf("starting server on %s", cfg.addr))

	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
