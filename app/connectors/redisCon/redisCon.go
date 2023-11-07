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

}

func GetRedisInstance() *redis.Client {
	return rdb
}
