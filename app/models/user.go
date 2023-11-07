package models

type User struct {
	ID                string `json:"id,omitempty"`
	Email             string `json:"email"`
	UserPassword      string `json:"user_password"`
	Gender            string `json:"gender"`
	About             string `json:"about"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
	AddressComplement string `json:"address_complement"`
	PostalCode        string `json:"postal_code"`
	City              string `json:"city"`
	Country           string `json:"country"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	IsStudent         bool   `json:"is_student"`
	IsHost            bool   `json:"is_host"`
}

type Users []User

// TableName is name of table or path for redis db
func (u User) TableName() string {
	return "user"
}
