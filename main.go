package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/easbarba/qas/internal/config"
)

var port = ":4000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/api/configs", configurations)

	log.Println(fmt.Sprintf("Server listening on %s", port))
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	// ignore all other routes other than root
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello, world!"))
}

// Return all configuration as JSON
func configurations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(config.AllToJson())
}
