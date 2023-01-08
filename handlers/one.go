package handlers

import (
	"net/http"

	"github.com/easbarba/qas_api/services"
)

// Return all configuration as JSON
func GetOneHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("lang")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(services.GetOne(key))
}
