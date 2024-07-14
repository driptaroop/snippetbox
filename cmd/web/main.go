package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"net/http"
	"os"
	"snippetbox.dripto.org/internal/models"
)

type Config struct {
	addr      string
	staticDir string
	dsn       string
}

func main() {
	//command line flags
	var config Config
	flag.StringVar(&config.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&config.staticDir, "staticDir", "./ui/static/", "static directory")
	flag.StringVar(&config.dsn, "dsn", "web:password@/snippetbox?parseTime=true", "Mysql data source name")
	flag.Parse()

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	//DB connections
	db, err := openDB(config.dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(db)

	//init template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// common config
	application := &Application{
		logger: logger,
		snippets: &models.SnippetModel{
			DB: db,
		},
		templateCache: templateCache,
	}

	// server start
	logger.Info(fmt.Sprintf("starting server on %s", config.addr))
	err = http.ListenAndServe(config.addr, application.routes(config.staticDir))
	logger.Error(err.Error())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
