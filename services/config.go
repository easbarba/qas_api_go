package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/easbarba/qas_api/common"
	"github.com/easbarba/qas_api/models"
)

func Append(project []models.Project) []models.Project {
	new := struct {
		Name   string "json:\"name\""
		Branch string "json:\"branch\""
		URL    string "json:\"url\""
	}{
		Name: "httprouter", Branch: "master", URL: "https://github.com/julienschmidt/httprouter",
	}

	return append(project, new)
}

// All configuration files unmarshallowed
func All() []models.Config {
	var configs []models.Config

	files, err := files()
	if err != nil {
		fmt.Println("no configuration file found!")
		os.Exit(1)
	}

	for _, file := range files {
		p := path.Join(common.QasConfigfolder, file.Name())
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

		configed := jsonToConfig(p)
		configs = append(configs, configed)
	}

	return configs
}

func GetOne(lang string) ([]byte, error) {
	for _, config := range All() {
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

func New() ([]byte, error) {
	config := models.Config{
		Lang: "elixir",
		Projects: []models.Project{
			{Name: "httprouter", Branch: "master", URL: "https://github.com/julienschmidt/httprouter"},
			{Name: "meh", Branch: "master", URL: "https://github.com/meh/meh"},
		},
	}

	writeNewConfig(config)

	result, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Write new configuration to a json file
func writeNewConfig(newConfig models.Config) error {
	configs := All()

	// Check if any configuration has already Lang set, and skip it!
	for _, config := range configs {
		if config.Lang == newConfig.Lang {
			log.Println("Configuration already exist. Skipping!")
			return nil
		}
	}

	// Write new configuration to file
	file, _ := json.MarshalIndent(newConfig, "", "  ")
	newConfigPath := path.Join(common.QasConfigfolder, newConfig.Lang+".json")
	err := ioutil.WriteFile(newConfigPath, file, 0644)
	if err != nil {
		return err
	}

	log.Println(fmt.Printf("%s configuration file saved on disk!", newConfig))

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
	for m, config := range configs {
		pjs, err := json.Marshal(config)
		if err != nil {
			log.Fatal("Error while marshalling configurations!")
		}

		result = append(result, pjs...)

		if m < len(configs)-1 {
			result = append(result, []byte(",")...)
		}
	}

	// append final right bracket
	result = append(result, []byte("]")...)

	return result
}

// Parse configuration file, check if the expect syntax is correct TODO: or err.
func jsonToConfig(filepath string) models.Config {
	var config models.Config

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

// all configuration files found TODO: return error if no configuration is found.
func files() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(common.QasConfigfolder)
	if err != nil {
		log.Fatal(err)
	}

	return files, nil
}
