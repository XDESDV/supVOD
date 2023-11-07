package models

type Kind struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type Kinds []Kind

// TableName is name of table or path for redis db
func (k Kind) TableName() string {
	return "kind"
}
