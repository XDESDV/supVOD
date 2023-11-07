package rediscon

import (
	"log"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func NewRedisClient() {
	log.Println("New redis client")
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Ping().Err()
	if err != nil {
		log.Fatalf("got nil, expected an error")
	} else {
		log.Println("REDIS successfully connected")
	}

}

func GetRedisInstance() *redis.Client {
	return rdb
}
