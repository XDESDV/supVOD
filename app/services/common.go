package services

import rediscon "supVOD/app/connectors/redisCon"

func listID(tablename string) ([]string, error) {
	var (
		val []string
		err error
	)

	redisInstance := rediscon.GetRedisInstance()

	key := tablename + "*"

	val, err = redisInstance.Keys(key).Result()
	if err == nil {
		return nil, err
	}

	return val, nil
}
