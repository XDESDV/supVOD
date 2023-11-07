package main

import (
	"log"
	"net/http"
	rediscon "supVOD/app/connectors/redisCon"
)

func main() {

	router := initialiseRouter()
	rediscon.NewRedisClient()
	log.Println("Listen & serve")
	http.ListenAndServe(":8080", router)

}
