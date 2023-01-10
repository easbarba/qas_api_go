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
	"log"
	"net/http"
	"os"

	"github.com/easbarba/qas_api/internal/controllers"
)

const (
	port = ":5000"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &controllers.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
	srv := &http.Server{
		Addr:     port,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	infoLog.Printf("Starting server on %s", port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
