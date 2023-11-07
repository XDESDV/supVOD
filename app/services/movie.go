package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	rediscon "supVOD/app/connectors/redisCon"
	redisconnector "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"

	"github.com/go-redis/redis"
)

func GetMovieByID(id string) (*models.Movie, error) {
	var (
		m   models.Movie
		err error
		val string
	)
	rdb := rediscon.GetRedisInstance()
	key := m.TableName() + "/" + id
	fmt.Println(key)
	if val, err = rdb.Get(key).Result(); err != redis.Nil {
		fmt.Println(val)
		if err = json.Unmarshal([]byte(val), &m); err == nil {
			return &m, err
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func CreateMovie(m models.Movie) error {
	var (
		err error
		val []byte
	)
	rdb := rediscon.GetRedisInstance()
	m.ID = functions.NewUUID()
	key := m.TableName() + "/" + m.ID
	if val, err = json.Marshal(m); err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}
	return err
}

func UpdateMovie(m models.Movie) error {
	var (
		cm  *models.Movie
		err error
		val []byte
	)
	if cm, err = GetMovieByID(m.ID); err == nil {
		if cm != nil {
			rdb := rediscon.GetRedisInstance()
			key := m.TableName() + "/" + m.ID
			if m.Title == "" {
				m.Title = cm.Title
			}
			if m.Description == "" {
				m.Description = cm.Description
			}
			if len(m.Kinds) == 0 {
				m.Kinds = cm.Kinds
			}
			if m.Duration == 0 {
				m.Duration = cm.Duration
			}
			fmt.Println("MOVIE : ", m)
			if val, err = json.Marshal(m); err == nil {
				err = rdb.Set(key, string(val), 0).Err()
			}
			return err
		} else {
			err = errors.New("id not exists")
			return err
		}
	} else {
		err = errors.New("error getbyid")
		return err
	}
}

func FindMovie(qm models.Query_Movie) (models.Movies, error) {
	var movies models.Movies
	var err error
	var m *models.Movie
	rdb := redisconnector.GetRedisInstance()
	if len(qm.IDs) == 0 && len(qm.Kinds) == 0 && len(qm.Titles) == 0 {
		for _, key := range rdb.Keys("*").Val() {
			table := strings.Split(key, "/")[0]
			id := strings.Split(key, "/")[1]
			if table == "movie" {
				if m, err = GetMovieByID(id); err == nil {
					movies = append(movies, *m)
				} else {
					return nil, err
				}
			}
		}
	} else {
		for _, key := range rdb.Keys("*").Val() {
			table := strings.Split(key, "/")[0]
			id := strings.Split(key, "/")[1]
			if table == "movie" {
				if m, err = GetMovieByID(id); err == nil {
					if len(qm.IDs) != 0 {
						for _, value := range qm.IDs {
							if m.ID == value {
								movies = append(movies, *m)
							}
						}
					}
					if len(qm.Kinds) != 0 {
						for _, value := range qm.IDs {
							if m.ID == value {
								movies = append(movies, *m)
							}
						}
					}
					if len(qm.Titles) != 0 {
						for _, value := range qm.IDs {
							if m.ID == value {
								movies = append(movies, *m)
							}
						}
					}
				} else {
					return nil, err
				}
			}
		}
	}
	return movies, nil
}
