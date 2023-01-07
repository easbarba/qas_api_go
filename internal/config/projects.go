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

type Project struct {
	Name   string `json:"name"`
	Branch string `json:"branch"`
	URL    string `json:"url"`
}

// Config structure of Configuration files
// log config files found
type Config struct {
	Lang     string    `json:"lang"`
	Projects []Project `json:"projects"`
}

func Append(project []Project) []Project {
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
func All() []Config {
	var configs []Config

	files, err := files()
	if err != nil {
		fmt.Println("no configuration file found!")
		os.Exit(1)
	}

	for _, file := range files {
		p := path.Join(QasConfigfolder, file.Name())
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

func New() ([]byte, error) {
	config := Config{
		Lang: "elixir",
		Projects: []Project{
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
func writeNewConfig(newConfig Config) error {
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
	newConfigPath := path.Join(QasConfigfolder, newConfig.Lang+".json")
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
	files, err := ioutil.ReadDir(QasConfigfolder)
	if err != nil {
		log.Fatal(err)
	}

	return files, nil
}
