package models

type Kind struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

func (k Kind) TableName() string {
	return "kinds"
}

func (k Kind) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":   k.ID,
		"name": k.Name,
	}
}

func (k Kind) Validate() bool {
	return k.Name != ""
}

func (k Kind) RequiredFieldsString() string {
	return "[name]"
}
