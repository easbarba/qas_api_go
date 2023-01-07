package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

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
