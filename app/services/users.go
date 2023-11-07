package services

import (
	"encoding/json"
	"log"
	rediscon "supVOD/app/connectors/redisConf"
	"supVOD/app/functions"
	"supVOD/app/models"
)

func CreateUser(user models.User) error {
	var (
		err error
		val []byte
	)
	redisInstance := rediscon.GetRedisInstance()

	user.ID = functions.NewUUID()
	key := user.TableName() + "/" + user.ID

	log.Println(user.ID)

	if val, err = json.Marshal(user); err == nil {
		err = redisInstance.Set(key, string(val), 0).Err()
	}

	log.Println(err)
	return err
}

func GetByID(id string) (*models.User, error) { // Renvoie un pointeur de User
	var (
		user models.User
		val  string
		err  error
	)

	redisInstance := rediscon.GetRedisInstance()

	key := user.TableName() + "/" + id

	// 1ère version, ou on chaine l'opération et le if en une ligne
	if val, err = redisInstance.Get(key).Result(); err == nil {
		if err = json.Unmarshal([]byte(val), &user); err != nil { // Unmarshal met les donnéees de val dans user
			return nil, err
		}
	}

	// 2ème version, plus lisible, ou on sépare l'opération et le if en deux lignes différentes
	// val, err = redisInstance.Get(key).Result()
	// if err == nil {
	// 	err = json.Unmarshal([]byte(val), &user)
	// 	if err != nil {
	// 		return nil ,err
	// 	}
	// }

	return &user, nil
}

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

// func DeleteUser(string id) {
// 	redisInstance := rediscon.GetRedisInstance()
// 	key := user.TableName() + "/" + user.ID

// 	deleted, err := yourRedisClient().Del(key).Result()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return deleted
// }
