package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	application := &Application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})),
		config: &Config{},
	}

	flag.StringVar(&application.config.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&application.config.staticDir, "staticDir", "./ui/static/", "static directory")
	flag.StringVar(&application.config.dsn, "dsn", "web:password@/snippetbox?parseTime=true", "Mysql data source name")
	flag.Parse()

	db, err := openDB(application.config.dsn)
	if err != nil {
		application.logger.Error(err.Error())
		os.Exit(1)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			application.logger.Error(err.Error())
		}
	}(db)

	application.logger.Info(fmt.Sprintf("starting server on %s", application.config.addr))

	err = http.ListenAndServe(application.config.addr, application.routes())
	application.logger.Error(err.Error())
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
