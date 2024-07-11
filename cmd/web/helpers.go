package main

import (
	"net/http"
	"runtime/debug"
)

func (app *Application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method string = r.Method
		uri    string = r.URL.RequestURI()
		trace         = string(debug.Stack())
	)
	app.logger.Error(err.Error(), "method", method, "uri", uri, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
