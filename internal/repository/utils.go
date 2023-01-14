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
	"os"
	"path"

	"github.com/easbarba/qas_api/internal/common"
	"github.com/easbarba/qas_api/internal/models"
)

// Write new configuration to a json file
func writeNewConfig(newConfig models.Config) error {
	configs, err := All()
	if err != nil {
		return errors.New("no configuration file found!")
	}

	// Check if any configuration has already Lang set, and skip it!
	for _, config := range configs {
		if config.Lang == newConfig.Lang {
			return errors.New("Configuration already exist. Skipping!")
		}
	}

	// Write new configuration to file
	file, _ := json.Marshal(newConfig.Projects)

	qas_dir, err := common.QasConfigfolder()
	if err != nil {
		return err
	}

	newConfigPath := path.Join(qas_dir, newConfig.Lang+".json")
	err = os.WriteFile(newConfigPath, file, 0644)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func RemoveConfig(lang string) error {
	qas_dir, err := common.QasConfigfolder()
	if err != nil {
		return err
	}

	configPath := path.Join(qas_dir, lang+".json")
	err = os.Remove(configPath)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Check for duplicates in configuration files
func ConfigCheckDuplicates() {
	panic("unimplemented")
}

// Bundle configurations as a JSON array
func AllToJson() ([]byte, error) {
	// mapped := make(map[string]models.Projects)

	configs, err := All()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// for _, config := range configs {
	// 	mapped[config.Lang] = config.Projects
	// }

	result, err := json.Marshal(configs)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil
}

// Parse configuration file,
// TODO: check if the expect syntax is correct
// TODO: or err.
func TranslateConfig(filepath string, filename string) (models.Config, error) {
	var projects models.Projects

	file, err := os.ReadFile(filepath)
	if err != nil {
		return models.Config{}, err
	}

	err = json.Unmarshal(file, &projects)
	if err != nil {
		errMsg := fmt.Sprintf("Configuration file has incorrect syntax \n%s", err.Error())
		return models.Config{}, errors.New(errMsg)
	}

	config := models.Config{
		Lang:     common.FileNameWithoutExtension(filename),
		Projects: projects,
	}

	return config, nil
}

func CheckConfigSyntax() error {
	panic("Not implemented")
}
