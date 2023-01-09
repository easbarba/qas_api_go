/*
*  Qas is free software: you can redistribute it and/or modify
*  it under the terms of the GNU General Public License as published by
*  the Free Software Foundation, either version 3 of the License, or
*  (at your option) any later version.

*  Qas is distributed in the hope that it will be useful,
*  but WITHOUT ANY WARRANTY; without even the implied warranty of
*  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*  GNU General Public License for more details.

*  You should have received a copy of the GNU General Public License
*  along with Qas. If not, see <https://www.gnu.org/licenses/>.
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"

	"github.com/easbarba/qas_api/internal/services"
)

const (
	port    = ":5000"
	version = "/v1"
	prefix  = "config"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}
	srv := &http.Server{
		Addr:     port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.index)
	mux.HandleFunc(app.routePath("all"), app.all)
	mux.HandleFunc(app.routePath("one"), app.one)
	mux.HandleFunc(app.routePath("create"), app.create)
	mux.HandleFunc(app.routePath("delete"), app.delete)

	return mux
}

func (app *application) routePath(resource string) string {
	return path.Join(version, prefix, resource)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// Welcome user
func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// ignore all other routes other than root
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	hello, err := json.Marshal("Hello, world!")
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(hello))
}

// Return all configuration as JSON
func (app *application) all(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(services.AllToJson())
}

// Return all configuration as JSON
func (app *application) one(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("lang")

	config, err := services.GetOne(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(config)
}

// Create a new configuration file
func (app *application) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	payload := r.Body
	defer r.Body.Close()

	newConfig, err := services.New(payload)
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(newConfig)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("lang")
	err := services.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		app.serverError(w, err)
	}

	w.WriteHeader(http.StatusNoContent)
}
