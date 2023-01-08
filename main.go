package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/easbarba/qas_api/handlers"
)

const (
	port    = ":4000"
	version = "/v1/"
)

func main() {
	routeList()

	log.Println(fmt.Sprintf("Server listening on %s", port))
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}

func routeList() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc(version+"configs", handlers.ListHandler)
	http.HandleFunc(version+"configs/view", handlers.GetOneHandler)
	http.HandleFunc(version+"configs/create", handlers.CreateHandler)
}
