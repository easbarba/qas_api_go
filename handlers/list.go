package handlers

import (
	"net/http"

	"github.com/easbarba/qas_api/internal/config"
)

// Return all configuration as JSON
func ListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(config.AllToJson())
}
