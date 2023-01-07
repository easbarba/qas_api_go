package handlers

import (
	"log"
	"net/http"

	"github.com/easbarba/qas_api/internal/config"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	newConfig, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newConfig)
}
