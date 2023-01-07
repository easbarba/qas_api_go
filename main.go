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
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/api/configs", configurations)

	log.Println(fmt.Sprintf("Server listening on %s", port))
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Return all configuration as json
func configurations(w http.ResponseWriter, r *http.Request) {
	configs := config.All()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("["))

	for m, config := range configs {
		pjs, err := json.Marshal(config)

		if err != nil {
			log.Fatal("no configuration file found!")
		}

		w.Write(pjs)

		if m < len(configs)-1 {
			w.Write([]byte(","))
		}
	}

	w.Write([]byte("]"))
}
