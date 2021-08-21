package cache

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

type RedisNil redis.Nil

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	if _, err := RedisClient.Ping().Result(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func Get(key string) *[]byte {
	val, err := RedisClient.Get(key).Bytes()
	if err == redis.Nil || err == nil {
		return &val
	}

	return nil
}
