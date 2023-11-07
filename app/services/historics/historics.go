package historics

import (
	"encoding/json"
	redisconnector "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"
)

func Find(userID string) ([]models.Historics, error) {
	var (
		err error
		val string
	)

	rdb := redisconnector.GetRedisInstance()

	var historics []models.Historics

	tmpHistoric := models.Historics{}
	keys, err := rdb.Keys(tmpHistoric.TableName() + "/*").Result()
	if err != nil {
		return historics, err
	}

	for _, key := range keys {
		val, err = rdb.Get(key).Result()
		if err == nil {
			var historic models.Historics
			err = json.Unmarshal([]byte(val), &historic)
			if err != nil {
				return historics, err
			}

			if userID == "" || historic.UserID == userID {
				historics = append(historics, historic)
			}
		}
	}

	return historics, err
}

func Create(historic *models.Historics) error {
	var (
		err error
		val []byte
	)

	rdb := redisconnector.GetRedisInstance()

	historic.ID = functions.NewUUID()
	key := historic.TableName() + "/" + historic.UserID + "/" + historic.MovieID
	key2 := historic.TableName() + "-id/" + historic.ID

	val, err = json.Marshal(historic)

	if err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}

	if err == nil {
		err = rdb.Set(key2, string(val), 0).Err()
	}

	return err
}

func GetById(id string) (*models.Historics, error) {
	var (
		historic models.Historics
		err      error
		val      string
	)

	rdb := redisconnector.GetRedisInstance()
	key := historic.TableName() + "-id/" + id

	if val, err = rdb.Get(key).Result(); err == nil {
		err = json.Unmarshal([]byte(val), &historic)
	}

	return &historic, err
}
