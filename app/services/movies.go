package services

import (
	"encoding/json"
	"log"
	rediscon "supVOD/app/connectors/redisConf"
	"supVOD/app/functions"
	"supVOD/app/models"
)

func CreateMovie(m models.Movie) error {
	var (
		err error
		val []byte
	)

	redisInstance := rediscon.GetRedisInstance()

	m.ID = functions.NewUUID()
	key := m.TableName() + "/" + m.ID

	log.Println(m.ID)

	if val, err = json.Marshal(m); err == nil {
		err = redisInstance.Set(key, string(val), 0).Err()
	}

	log.Println(err)
	return err
}

func GetMovieByID(id string) (*models.Movie, error) { // Renvoie un pointeur de movie
	var (
		movie models.Movie
		val   string
		err   error
	)

	redisInstance := rediscon.GetRedisInstance()

	key := movie.TableName() + "/" + id

	if val, err = redisInstance.Get(key).Result(); err == nil {
		if err = json.Unmarshal([]byte(val), &movie); err != nil { // Unmarshal met les donnéees de val dans movie
			return nil, err
		}
	}

	return &movie, nil
}

func GetMovies() (*models.Movies, error) { // Renvoie un pointeur de movie
	var (
		movie   models.Movie
		movies  models.Movies
		listVal []string
		val     string
		err     error
	)

	redisInstance := rediscon.GetRedisInstance()

	key := movie.TableName() + "*"
	log.Println(key)

	listVal, err = redisInstance.Keys(key).Result()
	if err != nil {
		return nil, err
	}

	for _, key := range listVal {
		if val, err = redisInstance.Get(key).Result(); err == nil {
			if err = json.Unmarshal([]byte(val), &movie); err != nil { // Unmarshal met les donnéees de val dans movie
				return nil, err
			}
			movies = append(movies, movie)
		}
	}
	return &movies, nil
}

func UpdateMovie(movie models.Movie) error {
	var (
		err error
		val []byte
	)

	redisInstance := rediscon.GetRedisInstance()
	key := movie.TableName() + "/" + movie.ID

	log.Println(key)

	if val, err = json.Marshal(movie); err == nil {
		log.Println(string(val))
		err = redisInstance.Set(key, string(val), 0).Err()
	}
	log.Println(err)

	return err
}

// func GetMoviesByTitle() (*models.Movies, error) { // TODO
// 	var (
// 		movie   models.Movie
// 		movies  models.Movies
// 		listVal []string
// 		val     string
// 		err     error
// 	)

// 	redisInstance := rediscon.GetRedisInstance()

// 	key := movie.TableName() + "*"
// 	log.Println(key)

// 	listVal, err = redisInstance.Keys(key).Result()
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, key := range listVal {
// 		if val, err = redisInstance.Get(key).Result(); err == nil {
// 			if err = json.Unmarshal([]byte(val), &movie); err == nil { // Unmarshal met les donnéees de val dans movie
// 				return nil, err
// 			}
// 			movies = append(movies, movie)
// 		}
// 	}
// 	return &movies, nil
// }
