package cache

import (
	"fmt"
	"go-image/config"
	"go-image/convert"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

var IsCache bool

func init() {

	IsCache = convert.StringToBool(config.GetSetting("redis.cache"))

	if IsCache {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.GetSetting("redis.addr"),
			Password: config.GetSetting("redis.password"),
			DB:       convert.StringToInt(config.GetSetting("redis.db")),
		})

		if _, err := redisClient.Ping().Result(); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}

}

func Set(key string, value interface{}, second int) {
	err := redisClient.Set(key, value, time.Second*time.Duration(second)).Err()
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
