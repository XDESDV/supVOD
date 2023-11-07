package services

import (
	"encoding/json"
	"strings"
	rediscon "supVOD/app/connectors/redisCon"
	redisconnector "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"

	"github.com/go-redis/redis"
)

func GetKindByID(id string) (*models.Kind, error) {
	var (
		k   models.Kind
		err error
		val string
	)
	rdb := rediscon.GetRedisInstance()
	key := k.TableName() + "/" + id
	if val, err = rdb.Get(key).Result(); err != redis.Nil {
		if err = json.Unmarshal([]byte(val), &k); err == nil {
			return &k, err
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func CreateKind(k models.Kind) error {
	var (
		err error
		val []byte
	)
	rdb := rediscon.GetRedisInstance()
	k.ID = functions.NewUUID()
	key := k.TableName() + "/" + k.ID
	if val, err = json.Marshal(k); err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}
	return err
}

func FindKind(qk models.Query_Kind) (models.Kinds, error) {
	var kinds models.Kinds
	var err error
	var k *models.Kind
	rdb := redisconnector.GetRedisInstance()
	if len(qk.IDs) == 0 && len(qk.Names) == 0 {
		for _, key := range rdb.Keys("*").Val() {
			table := strings.Split(key, "/")[0]
			id := strings.Split(key, "/")[1]
			if table == "kind" {
				if k, err = GetKindByID(id); err == nil {
					kinds = append(kinds, *k)
				} else {
					return nil, err
				}
			}
		}
	} else {
		for _, key := range rdb.Keys("*").Val() {
			table := strings.Split(key, "/")[0]
			id := strings.Split(key, "/")[1]
			if table == "kind" {
				if k, err = GetKindByID(id); err == nil {
					if len(qk.IDs) != 0 {
						for _, value := range qk.IDs {
							if k.ID == value {
								kinds = append(kinds, *k)
							}
						}
					}
					if len(qk.Names) != 0 {
						for _, value := range qk.Names {
							if k.Name == value {
								kinds = append(kinds, *k)
							}
						}
					}
				} else {
					return nil, err
				}
			}
		}
		var f_kinds models.Kinds
		for _, kind := range kinds {
			present := false
			for i := 0; i < len(f_kinds) && present == false; i++ {
				if kind.ID == f_kinds[i].ID {
					present = true
				}
			}
			if present == false {
				f_kinds = append(f_kinds, kind)
			}
		}
		kinds = f_kinds
	}
	return kinds, nil
}
