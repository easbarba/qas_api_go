package models

// Config structure of Configuration files
// log config files found
type Config struct {
	Lang     string    `json:"lang"`
	Projects []Project `json:"projects"`
}
