package controllers

import (
	"net/http"
	"path"
)

const (
	version = "/v1"
	prefix  = "config"
)

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.index)
	mux.HandleFunc(app.routePath("all"), app.all)
	mux.HandleFunc(app.routePath("one"), app.one)
	mux.HandleFunc(app.routePath("create"), app.create)
	mux.HandleFunc(app.routePath("delete"), app.delete)

	return mux
}

func (app *Application) routePath(resource string) string {
	return path.Join(version, prefix, resource)
}