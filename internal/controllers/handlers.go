package controllers

import (
	"net/http"

	"github.com/easbarba/qas_api/internal/repository"
)

// Return list configuration as JSON
func (app *Application) list(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result, err := repository.AllToJson()
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
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
func (app *Application) new(w http.ResponseWriter, r *http.Request) {
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

	app.InfoLog.Println("New configuration file saved on disk!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(newConfig)
}

func (app *Application) destroy(w http.ResponseWriter, r *http.Request) {
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
	app.InfoLog.Println("Configuration file deleted!")

}
