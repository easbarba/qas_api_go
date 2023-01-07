package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/easbarba/qas/internal/config"
)

var port = ":4000"

func main() {
	version := "/api/v1/"

	http.HandleFunc("/", indexHandler)
	http.HandleFunc(version+"configs", getHandler)

	log.Println(fmt.Sprintf("Server listening on %s", port))
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
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

// Return all configuration as JSON
func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(config.AllToJson())
}
