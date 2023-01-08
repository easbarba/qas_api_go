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
package common

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// ProjectsHomeFolder that all projects repositories will be stored at
var ProjectsHomeFolder string = path.Join(Home(), "Projects")

// QasConfigfolder that config files will be looked up for
var QasConfigfolder string = path.Join(Home(), ".config", "qas")

// Home folder of user
func Home() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return home
}

// all configuration files found TODO: return error if no configuration is found.
func Files() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(QasConfigfolder)
	if err != nil {
		log.Fatal(err)
	}

	return files, nil
}
