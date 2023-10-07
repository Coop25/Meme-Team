package models

type Meme struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Notes string   `json:"notes"`
	Links []string `json:"links"`
	Tags  []string `json:"tags"`

	// Note If this is empty it is not a local file and the first link should be used as the meme
	Path  string   `json:"path"`
}
