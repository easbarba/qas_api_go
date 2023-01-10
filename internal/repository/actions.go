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
package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/easbarba/qas_api/internal/common"
	"github.com/easbarba/qas_api/internal/models"
)

func Append(project models.Projects) models.Projects {
	panic("not implemented")
}

// All configuration files unmarshallowed
func All() ([]models.Config, error) {
	var configs []models.Config

	files, err := common.Files()
	if err != nil {
		return nil, errors.New("no configuration file found!")
	}

	for _, file := range files {
		qas_dir, err := common.QasConfigfolder()
		if err != nil {
			return nil, err
		}

		p := path.Join(qas_dir, file.Name())
		fileInfo, err := os.Stat(p)

		// ignore broken symbolic link
		if os.IsNotExist(err) {
			continue
		}

		// ignore directories
		if fileInfo.IsDir() {
			continue
		}

		// ignore csv files (legacy)
		if ext := filepath.Ext(p); ext == ".csv" {
			continue
		}

		configed, err := TranslateConfig(p, file.Name())
		if err != nil {
			return nil, err
		}

		configs = append(configs, configed)
	}

	return configs, nil
}

func GetOne(lang string) ([]byte, error) {
	configs, err := All()
	if err != nil {
		return nil, errors.New("no configuration file found!")
	}

	for _, config := range configs {
		if config.Lang == lang {
			cfg, err := json.Marshal(config)
			if err != nil {
				return nil, errors.New("Unable to convert current config to JSON")
			}

			return cfg, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("No configuration with Lang '%s' tag found!", lang))
}

func Delete(lang string) error {
	configs, err := All()
	if err != nil {
		return errors.New(err.Error())
	}

	for _, config := range configs {
		if config.Lang == lang {
			err := RemoveConfig(lang)
			if err != nil {
				return errors.New("Unable to delete config")
			}
		}

		return errors.New("No such a configuration file found!")
	}

	// successfully delete configuration file
	return nil
}

func New(payload io.ReadCloser) ([]byte, error) {
	var config models.Config
	err := json.NewDecoder(payload).Decode(&config)

	if err != nil {
		return nil, errors.New("jackshit")
	}

	err = writeNewConfig(config)
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	return result, nil
}
