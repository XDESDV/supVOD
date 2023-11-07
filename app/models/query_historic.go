package models

type Query_Historic struct {
	IDs       []string
	Users     Users
	Movies    Movies
	Durations []int
}
