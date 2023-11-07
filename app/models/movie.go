package models

type Movie struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	Kinds       Kinds  `json:"kinds,omitempty"`
}

type Movies []Movie

// TableName is name of table or path for redis db
func (m Movie) TableName() string {
	return "movie"
}
