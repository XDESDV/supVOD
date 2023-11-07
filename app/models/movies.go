package models

type Movie struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	Kinds       string `json:"kinds,omitempty"`
}

type Movies []Movie

func (m Movie) TableName() string {
	return "movie"
}
