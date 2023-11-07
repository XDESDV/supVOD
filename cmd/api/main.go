package main

import (
	"log"
	"net/http"
	rediscon "supVOD/app/connectors/redisConf"
	"supVOD/app/models"
	"supVOD/app/services"
)

func main() {
	rediscon.NewRedisClient()

	router := initialiseRouter()

	// addUser()
	// addMovie()

	log.Println("Server started on port 8080")

	http.ListenAndServe(":8080", router)
}

func addUser() {
	var u models.User
	u.FirstName = "John"
	u.LastName = "Doe"
	u.Email = "doe@gmail.com"
	u.UserPassword = "123456"

	if err := services.CreateUser(u); err != nil {
		log.Println(err)
	}
}

func addMovie() {
	var m models.Movie
	m.Title = "The Lord of the Rings: The Fellowship of the Ring"
	m.Description = "blablabla"
	m.Duration = 180

	if err := services.CreateMovie(m); err != nil {
		log.Println(err)
	}
}
