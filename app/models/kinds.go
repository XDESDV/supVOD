package models

type Kind struct {
	Name string
}

// type kinds []kind

type Kinds []Kind

func (u Kind) TableName() string {
	return "users"
}
