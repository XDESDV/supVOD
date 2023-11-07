package services

import (
	"encoding/json"
	"fmt"
	rediscon "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"
)

// ///////////////////////
// Create User
func CreateUser(user models.User) error {
	var (
		err error
		val []byte
	)

	redisInstance := rediscon.GetRedisInstance()
	//nouvel ID
	user.ID = functions.NewUUID()

	// clé dans mon table name qui définit la "connexion"
	key := user.TableName() + "/" + user.ID

	fmt.Println(user.ID)

	if val, err = json.Marshal(user); err == nil {
		err = redisInstance.Set(key, string(val), 0).Err()
	}
	return err
}

// GetByID

func GetByID(id string) (*models.User, error) {
	var (
		user models.User
		val  string
		err  error
	)

	redisInstance := rediscon.GetRedisInstance()
	key := user.TableName() + "/" + id

	val, err = redisInstance.Get(key).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &user); err != nil { //
			return nil, err
		}
	}
	return &user, nil
}

// UpdateUser
func UpdateUser(user models.User) error {
	var (
		err error
		val []byte
	)

	redisInstance := rediscon.GetRedisInstance()
	key := user.TableName() + "/" + user.ID

	if val, err = json.Marshal(user); err == nil {
		err = redisInstance.Set(key, string(val), 0).Err()
	}
	return err

}
