package services

import (
	rediscon "supVOD/app/connectors/redisCon"
)

func listID(tablename string) ([]string, error) {
	var (
		val []string
		err error
	)

	redisInstance := rediscon.GetRedisInstance() // récupération de la connexion ouverte

	key := tablename + "*" // récupération de la clé

	val, err = redisInstance.Keys(key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func listIDmovie(tablename string) ([]string, error) {
	var (
		val []string
		err error
	)

	redisInstance := rediscon.GetRedisInstance() // récupération de la connexion ouverte
	key := tablename + "*"                       // récupération de la clé

	val, err = redisInstance.Keys(key).Result()
	if err != nil {
		return err
	}
	return nil
}
