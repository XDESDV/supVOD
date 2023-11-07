package services

import (
	"encoding/json"
	"errors"
	rediscon "supVOD/app/connectors/redisCon"
	"supVOD/app/functions"
	"supVOD/app/models"

	"github.com/go-redis/redis"
)

func GetUserByID(id string) (*models.User, error) {
	var (
		u   models.User
		err error
		val string
	)
	rdb := rediscon.GetRedisInstance()
	key := u.TableName() + "/" + id
	if val, err = rdb.Get(key).Result(); err != redis.Nil {
		if err = json.Unmarshal([]byte(val), &u); err == nil {
			return &u, err
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func CreateUser(u models.User) error {
	var (
		err error
		val []byte
	)
	rdb := rediscon.GetRedisInstance()
	u.ID = functions.NewUUID()
	key := u.TableName() + "/" + u.ID
	if val, err = json.Marshal(u); err == nil {
		err = rdb.Set(key, string(val), 0).Err()
	}
	return err
}

func UpdateUser(u models.User) error {
	var (
		cu  *models.User
		err error
		val []byte
	)
	if cu, err = GetUserByID(u.ID); err == nil {
		if cu != nil {
			rdb := rediscon.GetRedisInstance()
			key := u.TableName() + "/" + u.ID
			if u.Email == "" {
				u.Email = cu.Email
			}
			if u.UserPassword == "" {
				u.UserPassword = cu.UserPassword
			}
			if u.Gender == "" {
				u.Gender = cu.Gender
			}
			if u.About == "" {
				u.About = cu.About
			}
			if u.Phone == "" {
				u.Phone = cu.Phone
			}
			if u.Address == "" {
				u.Address = cu.Address
			}
			if u.AddressComplement == "" {
				u.AddressComplement = cu.AddressComplement
			}
			if u.PostalCode == "" {
				u.PostalCode = cu.PostalCode
			}
			if u.City == "" {
				u.City = cu.City
			}
			if u.Country == "" {
				u.Country = cu.Country
			}
			if u.FirstName == "" {
				u.FirstName = cu.FirstName
			}
			if u.LastName == "" {
				u.LastName = cu.LastName
			}
			if val, err = json.Marshal(u); err == nil {
				err = rdb.Set(key, string(val), 0).Err()
			}
			return err
		} else {
			err = errors.New("id not exists")
			return err
		}
	} else {
		err = errors.New("error getbyid")
		return err
	}
}
