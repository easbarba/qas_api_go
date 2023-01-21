package controllers

import (
	"net/http"
	"path"

	"github.com/easbarba/qas_api/internal/middleware"
)

const (
	version  = "/v1"
	resource = "config"
)

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(app.routePath("list"), app.list)
	mux.HandleFunc(app.routePath("one"), app.one)
	mux.HandleFunc(app.routePath("new"), app.new)
	mux.HandleFunc(app.routePath("destroy"), app.destroy)

	return app.recoverPanic(
		app.logRequest(
			middleware.SecureHeaders(
				mux)))
}

func (app *Application) routePath(action string) string {
	return path.Join(version, resource, action)
}
