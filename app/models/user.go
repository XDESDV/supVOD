package models

type User struct {
	ID                string `json:"id,omitempty"`
	Email             string `json:"email"`
	UserPassword      string `json:"user_password"`
	Gender            string `json:"gender,omitempty"`
	About             string `json:"about,omitempty"`
	Phone             string `json:"phone,omitempty"`
	Address           string `json:"address,omitempty"`
	AddressComplement string `json:"address_complement,omitempty"`
	PostalCode        string `json:"postal_code,omitempty"`
	City              string `json:"city,omitempty"`
	Country           string `json:"country,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
}

type Users []User

// TableName is name of table or path for redis db
func (u User) TableName() string {
	return "user"
}
