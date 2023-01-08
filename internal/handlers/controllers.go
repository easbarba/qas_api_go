package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/easbarba/qas_api/internal/services"
)

// Welcome user
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// ignore all other routes other than root
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	hello, err := json.Marshal("Hello, world!")
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(hello))
}

// Create a new configuration file
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	payload := r.Body
	defer r.Body.Close()

	newConfig, err := services.New(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newConfig)
}

// Return all configuration as JSON
func ListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(services.AllToJson())
}

// Return all configuration as JSON
func GetOneHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("lang")

	config, err := services.GetOne(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(config)
}
