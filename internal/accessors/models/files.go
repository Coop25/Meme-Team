package models

type File struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Notes string   `json:"notes"`
	Links []string `json:"links"`
	Tags  []string `json:"tags"`
	Data  []byte   `json:"data"`
}
