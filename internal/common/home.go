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
	"os"
	"path"
)

// ProjectsHomeFolder that all projects repositories will be stored at
func ProjectsHomeFolder() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	return path.Join(home, "Projects"), nil
}

// QasConfigfolder that config files will be looked up for
func QasConfigfolder() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".config", "qas"), nil
}

// Home folder of user
func Home() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home, nil
}

// all configuration files found TODO: return error if no configuration is found.
func Files() ([]fs.FileInfo, error) {
	qas_folder, err := QasConfigfolder()
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(qas_folder)
	if err != nil {
		return nil, err
	}

	return files, nil
}
