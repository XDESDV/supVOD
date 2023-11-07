package models

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"first name"`
	LastName     string `json:"last name"`
	UserPassword string `json:"user password"`
	isAdmin      bool   `json:"-"`
}

type Users []User

func (u User) TableName() string {
	return "users"
}
