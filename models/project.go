package models

type Project struct {
	Name   string `json:"name"`
	Branch string `json:"branch"`
	URL    string `json:"url"`
}
