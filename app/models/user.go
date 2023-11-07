package models

type Users []User

type User struct {
	ID           string `json:"id,omitempty"`
	Email        string `json:"email"`
	UserPassword string `json:"user_password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

func (u *User) Validate() bool {
	return u.Email != "" && u.UserPassword != "" && u.FirstName != "" && u.LastName != ""
}

func (u User) TableName() string {
	return "users"
}

func (u User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"email":      u.Email,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
	}
}

func (u User) RequiredFieldsString() string {
	return "[email, first_name, last_name, user_password]"
}
