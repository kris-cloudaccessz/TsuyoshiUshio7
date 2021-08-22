package cache

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

const expireTime = time.Second * 600

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func Set(key string, value interface{}) {
	err := redisClient.Set(key, value, expireTime).Err()
	if err != nil {
		log.Println(err)
	}
}

func Get(key string) *[]byte {
	val, err := redisClient.Get(key).Bytes()
	if err == redis.Nil || err == nil {
		return &val
	}

	return nil
}
