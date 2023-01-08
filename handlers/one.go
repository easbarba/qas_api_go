package handlers

import (
	"log"
	"net/http"

	"github.com/easbarba/qas_api/services"
)

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
