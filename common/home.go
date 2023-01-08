package common

import (
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
