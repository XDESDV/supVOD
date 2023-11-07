package services

import (
	"encoding/json"
	"fmt"
	rediscon "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"
)

func CreateMovie(movie models.Movie) error {
	var (
		err error
		val []byte
	)

	redisInstance := rediscon.GetRedisInstance()
	//nouvel ID pour le Film
	movie.ID = functions.NewUUID()

	// clé dans mon table name qui définit la "connexion"
	key := movie.TableName() + "/" + movie.ID

	fmt.Println(movie.ID)

	if val, err = json.Marshal(movie); err == nil {
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

	redisInstance := rediscon.GetRedisInstance() // récupération de la connexion ouverte
	key := movie.TableName() + "/" + id          // récupération de la clé

	val, err = redisInstance.Get(key).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &movie); err != nil { //
			return nil, err
		}
	}
	return &movie, nil
}

func ListMovie(id string) (error, []models.Movie) {
	var (
		val []string
		err error
	)

	redisInstance := rediscon.GetRedisInstance() // récupération de la connexion ouverte
	key := movie.TableName() + "*"               // récupération de la clé

	val, err = redisInstance.Keys(key).Result()
	if err != nil {
		return nil
	}
	fmt.Println(val)
	return nil
}

func ListMovies() (models.Movie, error) {
	var (
		movie  models.Movie
		movies models.Movies
		val    string
		err    error
		listid []string
	)

	redisInstance := rediscon.GetRedisInstance() // récupération de la connexion ouverte

	listid, err = listIDmovie(movie.Tablename())
	if err != nil {
		return nil, err
	}

	for i := range listid {
		val, err = redisInstance.Get(listid[i]).Result()
		if err == nil {
			if err = json.Unmarshal([]byte(val), &movie); err != nil {
				return nil, err
			}
			users = append(movies, movie)
		}
	}
	return movies, nil
}
