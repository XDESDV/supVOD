package kinds

import (
	"encoding/json"
	redisconnector "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"
)

func Create(kind *models.Kind) error {
	var (
		err error
		val []byte
	)

	rdb := redisconnector.GetRedisInstance()

	kind.ID = functions.NewUUID()
	key := kind.TableName() + "/" + kind.ID

	if val, err = json.Marshal(kind); err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}

	return err
}

func List() ([]models.Kind, error) {
	var (
		err error
		val string
	)

	rdb := redisconnector.GetRedisInstance()
	var kinds []models.Kind

	keys, err := rdb.Keys("kinds/*").Result()
	if err != nil {
		return kinds, err
	}

	for _, key := range keys {
		if val, err = rdb.Get(key).Result(); err == nil {
			var kind models.Kind
			err = json.Unmarshal([]byte(val), &kind)
			if err != nil {
				return kinds, err
			}
			kinds = append(kinds, kind)
		}
	}

	return kinds, err
}

func GetById(id string) (*models.Kind, error) {
	var (
		kind models.Kind
		err  error
		val  string
	)

	rdb := redisconnector.GetRedisInstance()
	key := kind.TableName() + "/" + id

	val, err = rdb.Get(key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(val), &kind)
	}

	return &kind, err
}

func Update(kind *models.Kind) error {
	var (
		err error
		val []byte
	)

	rdb := redisconnector.GetRedisInstance()
	key := kind.TableName() + "/" + kind.ID

	if val, err = json.Marshal(kind); err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}

	return err
}
