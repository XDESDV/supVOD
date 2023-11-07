package models

type Movies struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"Duration"`
	Kinds       string `json:"kinds"`
}

type Movies []Movie

func (u Movie) TableName() string {
	return "movies"
}
