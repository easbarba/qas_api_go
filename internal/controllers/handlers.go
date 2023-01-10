package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/easbarba/qas_api/internal/repository"
)

// Welcome user
func (app *Application) index(w http.ResponseWriter, r *http.Request) {
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
func (app *Application) all(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(repository.AllToJson())
}

// Return all configuration as JSON
func (app *Application) one(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("lang")

	config, err := repository.GetOne(key)
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
func (app *Application) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	payload := r.Body
	defer r.Body.Close()

	newConfig, err := repository.New(payload)
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(newConfig)
}

func (app *Application) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("lang")
	err := repository.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		app.serverError(w, err)
	}

	w.WriteHeader(http.StatusNoContent)
}
