package services

import (
	rediscon "supVOD/app/connectors/redisConf"
	"supVOD/app/functions"
	"supVOD/app/models"
)

func CreateMovie(m models.Movie) {
	redisInstance := rediscon.GetRedisInstance()

	m.ID = functions.NewUUID()
	key := m.TableName() + "/" + m.ID
}
