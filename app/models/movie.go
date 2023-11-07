package models

type Movie struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Kinds       []Kind `json:"kinds"`
}

func (m Movie) TableName() string {
	return "movies"
}

func (m Movie) Validate() bool {
	return m.Title != "" && m.Description != "" && m.Duration >= 0
}

func (m Movie) RequiredFieldsString() string {
	return "[title, description, duration]"
}

func (m Movie) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          m.ID,
		"title":       m.Title,
		"description": m.Description,
		"duration":    m.Duration,
		"kinds":       m.Kinds,
	}
}
