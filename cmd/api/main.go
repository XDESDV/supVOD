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
	//addUser()

	log.Println("Listen & serve")
	http.ListenAndServe(":8080", router)

}

func addUser() {
	var u models.User

	u.FirstName = "Alexandre"
	u.LastName = "Huynh"
	u.Email = "alexandre.huynh456@icloud.com"
	u.UserPassword = "123456"

	if err := services.CreateUser(u); err != nil {
		fmt.Println(err)

	}
}
