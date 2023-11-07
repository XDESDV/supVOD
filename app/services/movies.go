package services

import (
	"encoding/json"
	"fmt"
	rediscon "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"
)

// CreateMovie takes a movie model and stores it in Redis.
func CreateMovie(movie models.Movie) error {
	var (
		err error
		val []byte
	)

	redisInstance := rediscon.GetRedisInstance()
	// Generate a new UUID for the movie.
	movie.ID = functions.NewUUID()

	// Define the key for the movie using the table name and ID.
	key := movie.TableName() + "/" + movie.ID

	fmt.Println(movie.ID)

	// Marshal the movie data into JSON.
	if val, err = json.Marshal(movie); err == nil {
		// Set the value in Redis using the key.
		err = redisInstance.Set(key, string(val), 0).Err()
	}
	return err
}

func GetMovieByID(id string) (*models.Movie, error) {
	var (
		movie models.Movie
		val   string
		err   error
	)

	redisInstance := rediscon.GetRedisInstance()
	key := movie.TableName() + "/" + id

	val, err = redisInstance.Get(key).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &movie); err != nil {
			return nil, err
		}
	}
	return &movie, nil
}

func UpdateMovie(movie models.Movie) error {
	var (
		err error
		val []byte
	)

	redisInstance := rediscon.GetRedisInstance()
	key := movie.TableName() + "/" + movie.ID

	if val, err = json.Marshal(movie); err == nil {
		err = redisInstance.Set(key, string(val), 0).Err()
	}
	return err
}
