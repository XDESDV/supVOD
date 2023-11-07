package services

import (
	"encoding/json"
	"fmt"
	rediscon "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"
)

// ///////////////////////
// Create Movie
func CreateMovie(movie models.Movie) error {
	var (
		err error
		val []byte
	)

	redisInstance := rediscon.GetRedisInstance()
	//nouveau movie
	movie.ID = functions.NewUUID()

	// clé dans le table name qui définit la "connexion"
	key := movie.TableName() + "/" + movie.ID

	fmt.Println(movie.ID)

	if val, err = json.Marshal(movie); err == nil {
		err = redisInstance.Set(key, string(val), 0).Err()
	}
	return err
}

// ////////////////
// GetByIDMovie
func GetByIDMovie(id string) (*models.Movie, error) {
	var (
		movie models.Movie
		val   string
		err   error
	)

	redisInstance := rediscon.GetRedisInstance()
	key := "movies/" + id

	val, err = redisInstance.Get(key).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &movie); err != nil {
			return nil, err
		}
	}
	return &movie, nil
}

// ////////////////
// UpdateMovie
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

//////////////////
//listMovies

func ListMovies(movie models.Movie) error {
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
