package models

type MovieListResult struct {
	Movies     []Movie
	TotalCount int64
	MaxPage    int
}
