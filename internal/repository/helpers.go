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
	"log"
	"os"
	"path"

	"github.com/easbarba/qas_api/internal/common"
	"github.com/easbarba/qas_api/internal/models"
)

// Write new configuration to a json file
func writeNewConfig(newConfig models.Config) error {
	configs := All()

	// Check if any configuration has already Lang set, and skip it!
	for _, config := range configs {
		if config.Lang == newConfig.Lang {
			return errors.New("Configuration already exist. Skipping!")
		}
	}

	// Write new configuration to file
	file, _ := json.MarshalIndent(newConfig.Projects, "", "  ")
	newConfigPath := path.Join(common.QasConfigfolder, newConfig.Lang+".json")
	err := os.WriteFile(newConfigPath, file, 0644)
	if err != nil {
		return errors.New(err.Error())
	}

	log.Println(fmt.Printf("%s configuration file saved on disk!", newConfig))
	return nil
}

func RemoveConfig(lang string) error {
	configPath := path.Join(common.QasConfigfolder, lang+".json")

	err := os.Remove(configPath)
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
func AllToJson() []byte {
	// begin json object with a left bracket
	result := []byte("[")

	// append a colon to each object configuration
	configs := All()
	for index, config := range configs {
		projects, err := json.Marshal(config)
		if err != nil {
			log.Println("Error while marshalling configurations!")
		}

		result = append(result, projects...)

		if index < len(configs)-1 {
			result = append(result, []byte(",")...)
		}
	}

	// append final right bracket
	result = append(result, []byte("]")...)

	return result
}

// Parse configuration file,
// TODO: check if the expect syntax is correct
// TODO: or err.
func TranslateConfig(filepath string, filename string) models.Config {
	var projects models.Projects

	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(file, &projects)
	if err != nil {
		errMsg := fmt.Sprintf("Configuration file has incorrect syntax \n%s", err)
		log.Println(errMsg)
	}

	config := models.Config{
		Lang:     common.FileNameWithoutExtension(filename),
		Projects: projects,
	}

	return config
}

func CheckConfigSyntax() error {
	panic("Not implemented")
}
