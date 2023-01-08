package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/easbarba/qas_api/handlers"
)

var port = ":4000"

func main() {
	version := "/api/v1/"

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc(version+"configs", handlers.ListHandler)
	http.HandleFunc(version+"configs?lang", handlers.GetOneHandler)
	http.HandleFunc(version+"configs/create", handlers.CreateHandler)

	log.Println(fmt.Sprintf("Server listening on %s", port))
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}
