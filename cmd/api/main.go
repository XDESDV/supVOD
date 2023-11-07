package main

import (
	"github.com/joho/godotenv"
	"log"
	rediscon "supVOD/app/connectors/redisCon"
	"supVOD/app/routers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		panic("INIT_PANIC: Error loading .env file")
	}

	err = rediscon.NewRedisClient()
	if err != nil {
		log.Println(err)
		panic("INIT_PANIC: Redis failed to connect")
	}

	var server routers.Server
	server.Init()

	err = server.Run()
	if err != nil {
		log.Println(err)
		panic("INIT_PANIC: Server failed to start")
	}

}
