package main

import (
	"net/http"
	"time"
)

func (app *application) serve() error {
	svr := http.Server{
		Addr:        app.addr,
		ReadTimeout: 2 * time.Second,
		Handler:     app.routes(),
	}
	return svr.ListenAndServe()
}
