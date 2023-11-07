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
