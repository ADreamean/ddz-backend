package user

import (
	"github.com/ADreamean/ddz-backend/until/redis"
	"log"
	"fmt"
)

type User struct {
	Id   int32
	Name string
}

func Create(name string) *User {
	id, err := redis.Int("INCR", "user")
	if err != nil {
		log.Panic(err)
	}

	user := User{int32(id), name}
	if _, err := redis.Do("SET", fmt.Sprintf("user_%d", id), user); err != nil {
		log.Panic(err)
	}

	return &user
}

func Find(id int32) *User {
	redis.Do("get", fmt.Sprintf("user_%d", id))
}
