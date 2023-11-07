package models

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Firstname    string `json:"fisrt_name"`
	Lastname     string `json:"last_name"`
	Userpassword string `json:"user_password"`
	isAdmin      bool   `json:"-"`
}

// type Users []person

type Users []User

func (u User) TableName() string {
	return "users"
}
