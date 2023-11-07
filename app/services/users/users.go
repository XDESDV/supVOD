package users

import (
	"encoding/json"
	redisconnector "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"
)

func GetUserByID(id string) (*models.User, error) {
	var (
		u   models.User
		err error
		val string
	)

	rdb := redisconnector.GetRedisInstance()
	key := u.TableName() + "/" + id

	if val, err = rdb.Get(key).Result(); err == nil {
		err = json.Unmarshal([]byte(val), &u)
	}

	return &u, err
}

func CreateUser(u models.User) (models.User, error) {
	var (
		err error
		val []byte
	)

	rdb := redisconnector.GetRedisInstance()

	u.ID = functions.NewUUID()
	key := u.TableName() + "/" + u.ID

	if val, err = json.Marshal(u); err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}

	return u, err
}

func UpdateUser(u models.User) error {
	var (
		err error
		val []byte
	)

	rdb := redisconnector.GetRedisInstance()

	key := u.TableName() + "/" + u.ID

	val, err = json.Marshal(u)
	if err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}

	return err
}
