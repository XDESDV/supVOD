package models

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserPassword string `json:"user_password"`
	isAdmin      bool   `json:"-"`
}

type Users []User

func (u User) TableName() string {
	return "users"
}
