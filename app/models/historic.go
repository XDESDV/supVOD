package models

type Historic struct {
	ID       string `json:"id,omitempty"`
	User     User   `json:"user"`
	Movie    Movie  `json:"movie"`
	Duration int    `json:"duration"`
}

type Historics []Historic

// TableName is name of table or path for redis db
func (h Historic) TableName() string {
	return "hictoric"
}
