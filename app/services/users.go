package services

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
		if err = json.Unmarshal([]byte(val), &u); err == nil {
			return nil, err
		}
	}
	return &u, nil
}

func CreateUser(u models.User) error {
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

	return err
}

func UpdateUser(u models.User) error {
	var (
		err error
		val []byte
	)

	rdb := redisconnector.GetRedisInstance()

	key := u.TableName() + "/" + u.ID

	if val, err = json.Marshal(u); err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}

	return err
}

func FindUsers(u models.User) {
	// var (
	// 	err error
	// )

	// rdb := redisconnector.GetRedisInstance()

	// Set some fields.
	// if _, err := rdb.Pipelined(func(rdb redis.Pipeliner) error {

	// 	rdb.HSet(key, "str1", "hello")

	// 	return nil
	// }); err != nil {
	// 	panic(err)
	// }

	// var users models.Users

	// // Scan all fields into the model.
	// if err := rdb.HGetAll(key).

}
