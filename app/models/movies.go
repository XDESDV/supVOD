package models

type Movie struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Duration    int     `json:"duration"`
	Kinds       []Kinds `json:"kinds"`
}

type Movies []Movie

func (m Movie) TableName() string {
	return "movie"
}
