package movies

import (
	"encoding/json"
	"math"
	"strings"
	redisconnector "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"
)

func GetMovieByID(id string) (*models.Movie, error) {
	var (
		movie models.Movie
		err   error
		val   string
	)

	rdb := redisconnector.GetRedisInstance()
	key := movie.TableName() + "/" + id

	if val, err = rdb.Get(key).Result(); err == nil {
		err = json.Unmarshal([]byte(val), &movie)
	}

	return &movie, err
}

func List(pagination models.PaginationParams, search models.SearchParams) (models.MovieListResult, error) {
	result := models.MovieListResult{
		Movies:     []models.Movie{},
		TotalCount: 0,
		MaxPage:    0,
	}

	rdb := redisconnector.GetRedisInstance()
	pattern := models.Movie{}.TableName() + "/*"

	var allMatchingMovies []models.Movie

	iter := rdb.Scan(0, pattern, 0).Iterator()
	for iter.Next() {
		var movie models.Movie
		val, err := rdb.Get(iter.Val()).Result()
		if err != nil {
			continue
		}
		err = json.Unmarshal([]byte(val), &movie)
		if err != nil {
			continue
		}

		if search.Search == "" && len(search.Kinds) == 0 {
			allMatchingMovies = append(allMatchingMovies, movie)
			continue
		}

		if search.Search != "" && strings.Contains(strings.ToLower(movie.Title), strings.ToLower(search.Search)) {
			allMatchingMovies = append(allMatchingMovies, movie)
			continue
		}

		if len(search.Kinds) > 0 {
		FIND:
			for _, kind := range movie.Kinds {
				for _, searchKind := range search.Kinds {
					if strings.ToLower(kind.Name) == strings.ToLower(searchKind.Name) {
						allMatchingMovies = append(allMatchingMovies, movie)
						break FIND
					}
				}
			}
		}
	}

	if err := iter.Err(); err != nil {
		return result, err
	}

	totalCount := len(allMatchingMovies)
	result.TotalCount = int64(totalCount)
	result.MaxPage = int(math.Ceil(float64(totalCount) / float64(pagination.PageSize)))

	start := (pagination.Page - 1) * pagination.PageSize
	end := start + pagination.PageSize
	if end > totalCount {
		end = totalCount
	}

	if start < totalCount {
		result.Movies = allMatchingMovies[start:end]
	}

	return result, nil
}

func Create(movie *models.Movie) error {
	var (
		err error
		val []byte
	)

	rdb := redisconnector.GetRedisInstance()

	movie.ID = functions.NewUUID()
	key := movie.TableName() + "/" + movie.ID

	if val, err = json.Marshal(movie); err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}

	return err
}

func UpdateMovie(movie *models.Movie) error {
	var (
		err error
		val []byte
	)

	rdb := redisconnector.GetRedisInstance()
	key := movie.TableName() + "/" + movie.ID

	if val, err = json.Marshal(movie); err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}

	return err
}
