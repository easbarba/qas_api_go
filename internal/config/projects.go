package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// HomeFolder that all projects repositories will be stored at
var HomeFolder string = path.Join(Home(), "Projects")

// folder that config files will be looked up for
var folder string = path.Join(Home(), ".config", "qas")

// Home folder of user
func Home() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return home
}

type Projects []struct {
	Name   string `json:"name"`
	Branch string `json:"branch"`
	URL    string `json:"url"`
}

// Config structure of Configuration files
// log config files found
type Config struct {
	Lang     string `json:"lang"`
	Projects `json:"projects"`
}

func Append(projects Projects) Projects {
	new := struct {
		Name   string "json:\"name\""
		Branch string "json:\"branch\""
		URL    string "json:\"url\""
	}{
		Name: "httprouter", Branch: "master", URL: "https://github.com/julienschmidt/httprouter",
	}

	return append(projects, new)
}

// All configuration files unmarshallowed
func All() []Config {
	var configs []Config

	files, err := files()
	if err != nil {
		fmt.Println("no configuration file found!")
		os.Exit(1)
	}

	for _, file := range files {
		p := path.Join(folder, file.Name())
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
func jsonToConfig(filepath string) Config {
	var config Config

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
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	return files, nil
}
