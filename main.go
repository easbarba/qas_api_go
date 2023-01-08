/*
*  Qas is free software: you can redistribute it and/or modify
*  it under the terms of the GNU General Public License as published by
*  the Free Software Foundation, either version 3 of the License, or
*  (at your option) any later version.

*  Qas is distributed in the hope that it will be useful,
*  but WITHOUT ANY WARRANTY; without even the implied warranty of
*  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*  GNU General Public License for more details.

*  You should have received a copy of the GNU General Public License
*  along with Qas. If not, see <https://www.gnu.org/licenses/>.
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/easbarba/qas_api/internal/handlers"
)

const (
	port    = ":5000"
	version = "/v1/"
)

func main() {
	routeList()

	msg := fmt.Sprintf("Server listening on %s", port)
	log.Println(msg)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func routeList() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc(version+"cfg/all", handlers.ListHandler)
	http.HandleFunc(version+"cfg/one", handlers.GetOneHandler)
	http.HandleFunc(version+"cfg/create", handlers.CreateHandler)
	http.HandleFunc(version+"cfg/delete", handlers.DeleteHandler)
}
