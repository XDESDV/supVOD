package main

import (
	"fmt"
	"log"
	"net/http"
	rediscon "supVOD/app/connectors/redisCon"
	"supVOD/app/models"
	"supVOD/app/services"
)

func main() {
	rediscon.NewRedisClient()

	router := initialiseRouter()
	addUser()
	addMovie()
	log.Println("Listen & serve")
	http.ListenAndServe(":8080", router)

}

func addUser() {
	var u models.User

	u.FirstName = "Benjamin"
	u.LastName = "Barillot"
	u.Email = "benmios@hotmail.fr"
	u.UserPassword = "123456"

	if err := services.CreateUser(u); err != nil {
		fmt.Println(err)
	}
}

func addMovie() {
	var u models.Movie

	u.Title = "SNK"
	u.Description = "Attack on Titan"
	u.Duration = 25
	u.Kinds = "Anime"

	if err := services.CreateMovie(u); err != nil {
		fmt.Println(err)
	}
}
