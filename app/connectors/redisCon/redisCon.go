package rediscon

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func NewRedisClient() error {
	log.Println("New redis client")
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	err := rdb.Ping().Err()
	if err != nil {
		return err
	}

	return nil
}

func GetRedisInstance() *redis.Client {
	return rdb
}
